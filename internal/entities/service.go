package entities

import (
	"context"

	"github.com/papey08/apcor/pkg/errstream"
)

type Service struct {
	*Repo

	preBuild Commands
	build    Commands
	run      *Command
	debug    *Command
	stop     *Command
}

func NewService(
	repoUrl string,
	path string,
	preBuild []string,
	build []string,
	run string,
	debug string,
	stop string,
) (*Service, error) {
	repo, err := NewRepo(repoUrl, path)
	if err != nil {
		return nil, err
	}
	return &Service{
		Repo:     repo,
		preBuild: NewCommands(preBuild...),
		build:    NewCommands(build...),
		run:      NewCommand(run),
		debug:    NewCommand(debug),
		stop:     NewCommand(stop),
	}, nil
}

func (s *Service) PreBuild(ctx context.Context, es *errstream.ErrStream) error {
	for _, cmd := range s.preBuild {
		cmd.ExecAsync(ctx, s.path, es)
	}
	return nil
}

func (s *Service) Build(ctx context.Context) error {
	for _, cmd := range s.build {
		cmd.Exec(ctx, s.path)
	}
	return nil
}

func (s *Service) Run(ctx context.Context, es *errstream.ErrStream) error {
	s.run.ExecAsync(ctx, s.path, es)
	return nil
}

func (s *Service) Debug(ctx context.Context, es *errstream.ErrStream) error {
	s.debug.ExecAsync(ctx, s.path, es)
	return nil
}

func (s *Service) Stop(ctx context.Context, es *errstream.ErrStream) (bool, error) {
	if s.stop == nil {
		return false, nil
	}
	s.stop.ExecAsync(ctx, s.path, es)
	return true, nil
}
