package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/papey08/apcor/internal/entities"
	"github.com/papey08/apcor/internal/errs"
	"github.com/papey08/apcor/internal/parser"
	"github.com/papey08/apcor/pkg/errstream"
	"github.com/papey08/apcor/pkg/utils"
)

type Manager struct {
	entities.Accessor

	InstallDeps map[string]entities.Commands

	PreTasks  entities.Commands
	PostTasks entities.Commands

	Libs     map[string]*entities.Lib
	Services map[string]*entities.Service

	stopCtx   map[string]context.CancelFunc
	ErrStream *errstream.ErrStream
}

func NewManager(data *parser.ApcorParsingData) (*Manager, error) {
	m := &Manager{}

	m.InstallDeps = make(map[string]entities.Commands)
	for osName, commands := range data.InstallDeps {
		m.InstallDeps[osName] = entities.NewCommands(commands...)
	}

	m.PreTasks = entities.NewCommands(data.Common.PreTasks...)
	m.PostTasks = entities.NewCommands(data.Common.PostTasks...)

	m.Libs = make(map[string]*entities.Lib)
	for name, lib := range data.Libs {
		var err error
		m.Libs[name], err = entities.NewLib(lib.RemoteRepoUrl, lib.Path, lib.Build)
		if err != nil {
			return nil, errors.Join(errs.InvalidLib, err)
		}
	}

	m.Services = make(map[string]*entities.Service)
	for name, service := range data.Services {
		var err error
		m.Services[name], err = entities.NewService(service.RemoteRepoUrl, service.Path, service.PreBuild, service.Build, service.Run, service.Debug, service.Stop)
		if err != nil {
			return nil, errors.Join(errs.InvalidService, err)
		}
	}

	m.stopCtx = make(map[string]context.CancelFunc)
	m.ErrStream = errstream.New(len(data.Libs) + len(data.Services))

	return m, nil
}

func (m *Manager) RunInstallDeps(ctx context.Context, platform string) {
	if commands, ok := m.InstallDeps[platform]; ok {
		for _, c := range commands {
			err := c.ExecInteractive(ctx, entities.DefaultPath)
			if err != nil {
				m.ErrStream.ErrChan <- err
			}
		}
	} else if platform != "" {
		m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnsupportedPlatform, platform)
	} else {
		m.ErrStream.ErrChan <- fmt.Errorf("%w: <empty string>", errs.UnsupportedPlatform)
	}
}

// ExecPreTasks executes every pre-task command asynchronously
func (m *Manager) ExecPreTasks(ctx context.Context) {
	for _, task := range m.PreTasks {
		task.ExecAsync(ctx, entities.DefaultPath, m.ErrStream)
	}
}

// ExecPostTasks executes every post-task command consistently
func (m *Manager) ExecPostTasks(ctx context.Context) {
	for _, task := range m.PostTasks {
		err := task.Exec(ctx, entities.DefaultPath)
		if err != nil {
			m.ErrStream.ErrChan <- err
		}
	}
}

// Clone clones repositories consistently
func (m *Manager) Clone(ctx context.Context, repos []string) {
	reposMap := utils.SliceToSet(repos)

	// clone libs
	for r := range reposMap {
		if lib, ok := m.Libs[r]; ok {
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := lib.Clone(ctx)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.CloneRepo, r, err.Error())
			}
		}
		delete(reposMap, r)
	}

	// clone services
	for r := range reposMap {
		if srv, ok := m.Services[r]; ok {
			m.cancel(r)
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := srv.Clone(ctx)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.CloneRepo, r, err.Error())
			}
		}
		delete(reposMap, r)
	}

	// unknown repos
	for r := range reposMap {
		m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnknownRepo, r)
	}
}

// Checkout checkouts repositories consistently
func (m *Manager) Checkout(ctx context.Context, repos []string, branch string) {
	reposMap := utils.SliceToSet(repos)

	// checkout libs
	for r := range reposMap {
		if lib, ok := m.Libs[r]; ok {
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := lib.Checkout(ctx, branch)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.CheckoutRepo, r, err.Error())
			}
		}
		delete(reposMap, r)
	}

	// checkout services
	for r := range reposMap {
		if srv, ok := m.Services[r]; ok {
			m.cancel(r)
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := srv.Checkout(ctx, branch)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.CheckoutRepo, r, err.Error())
			}
		}
		delete(reposMap, r)
	}

	// unknown repos
	for r := range reposMap {
		m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnknownRepo, r)
	}
}

// Branch returns map of repo name -> current branch of repo
func (m *Manager) Branch(ctx context.Context) map[string]string {
	res := make(map[string]string, len(m.Libs)+len(m.Services))

	var err error
	for name, lib := range m.Libs {
		res[name], err = lib.Branch(ctx)
		if err != nil {
			m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.BranchRepo, name, err.Error())
		}
	}

	for name, srv := range m.Services {
		res[name], err = srv.Branch(ctx)
		if err != nil {
			m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.BranchRepo, name, err.Error())
		}
	}

	return res
}

// PullBranch makes git pull in all given repos consistently
func (m *Manager) Pull(ctx context.Context, repos []string) {
	reposMap := utils.SliceToSet(repos)

	// build libs
	for r := range reposMap {
		if lib, ok := m.Libs[r]; ok {
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := lib.Pull(ctx)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.LibBuild, r)
			}
		}
		delete(reposMap, r)
	}

	// build services
	for r := range reposMap {
		if srv, ok := m.Services[r]; ok {
			m.cancel(r)
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := srv.Pull(ctx)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.ServiceBuild, r)
			}
		}
		delete(reposMap, r)
	}

	// unknown repos
	for r := range reposMap {
		m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnknownRepo, r)
	}
}

// PreBuild executes every pre-build command of services asynchronously
func (m *Manager) PreBuild(ctx context.Context, services []string) error {
	for _, srvName := range services {
		if srv, ok := m.Services[srvName]; ok {
			m.cancel(srvName)
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[srvName] = cancel
			srv.PreBuild(ctx, m.ErrStream)
		} else {
			m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnknownRepo, srvName)
		}
	}
	return nil
}

// Build executes every build command of libs, than of services consistently
func (m *Manager) Build(ctx context.Context, repos []string) error {
	reposMap := utils.SliceToSet(repos)

	// build libs
	for r := range reposMap {
		if lib, ok := m.Libs[r]; ok {
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := lib.Build(ctx)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.LibBuild, r, err.Error())
			}
		}
		delete(reposMap, r)
	}

	// build services
	for r := range reposMap {
		if srv, ok := m.Services[r]; ok {
			m.cancel(r)
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[r] = cancel
			err := srv.Build(ctx)
			if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.ServiceBuild, r, err.Error())
			}
		}
		delete(reposMap, r)
	}

	// unknown repos
	for r := range reposMap {
		m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnknownRepo, r)
	}

	return nil
}

// Run executes every run command of services asynchronously
func (m *Manager) Run(ctx context.Context, services []string) error {
	for _, srvName := range services {
		if srv, ok := m.Services[srvName]; ok {
			m.cancel(srvName)
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[srvName] = cancel
			srv.Run(ctx, m.ErrStream)
		} else {
			m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnknownRepo, srvName)
		}
	}
	return nil
}

// Debug executes every debug command of services asynchronously
func (m *Manager) Debug(ctx context.Context, services []string) error {
	for _, srvName := range services {
		if srv, ok := m.Services[srvName]; ok {
			ctx, cancel := context.WithCancel(ctx)
			m.stopCtx[srvName] = cancel
			srv.Debug(ctx, m.ErrStream)
		} else {
			m.ErrStream.ErrChan <- fmt.Errorf("%w: %s", errs.UnknownRepo, srvName)
		}
	}
	return nil
}

// Stop executes every stop command of services asynchronously, if service have no stop command, currently executing command will exit
func (m *Manager) Stop(ctx context.Context, services []string) error {
	for _, srvName := range services {
		if srv, ok := m.Services[srvName]; ok {
			if ok, err := srv.Stop(ctx, m.ErrStream); !ok {
				m.cancel(srvName)
			} else if err != nil {
				m.ErrStream.ErrChan <- fmt.Errorf("%w: %s: %s", errs.ServiceStop, srvName, err.Error())
			}
		}
	}
	return nil
}

func (m *Manager) cancel(serviceName string) {
	if cancel, ok := m.stopCtx[serviceName]; ok {
		cancel()
	}
}

func (m *Manager) updateCancel(ctx context.Context, serviceName string) context.Context {
	m.cancel(serviceName)
	ctx, cancel := context.WithCancel(ctx)
	m.stopCtx[serviceName] = cancel
	return ctx
}
