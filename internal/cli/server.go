package cli

import (
	"context"

	"github.com/papey08/apcor/internal/cli/gen"
	"github.com/papey08/apcor/internal/entities"
	"github.com/papey08/apcor/internal/manager"
	"google.golang.org/grpc"
)

//go:generate protoc cli_service.proto --proto_path=. --go-grpc_out=require_unimplemented_servers=false:. --go_out=.

type Server struct {
	mng *manager.Manager
}

var _ gen.CliServiceServer = (*Server)(nil)

func NewServer(mng *manager.Manager) *grpc.Server {
	s := grpc.NewServer()
	gen.RegisterCliServiceServer(s, &Server{
		mng: mng,
	})
	return s
}

func (s *Server) InstallDeps(ctx context.Context, req *gen.InstallDepsRequest) (*gen.InstallDepsResponse, error) {
	s.mng.RunInstallDeps(ctx, req.Platform)
	return &gen.InstallDepsResponse{}, nil
}

func (s *Server) ExecPreTasks(ctx context.Context, req *gen.ExecPreTasksRequest) (*gen.ExecPreTasksResponse, error) {
	s.mng.ExecPreTasks(ctx)
	return &gen.ExecPreTasksResponse{}, nil
}

func (s *Server) ExecPostTasks(ctx context.Context, req *gen.ExecPostTasksRequest) (*gen.ExecPostTasksResponse, error) {
	s.mng.ExecPostTasks(ctx)
	return &gen.ExecPostTasksResponse{}, nil
}

func (s *Server) Clone(ctx context.Context, req *gen.CloneRequest) (*gen.CloneResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.Clone(ctx, repos)
	return &gen.CloneResponse{}, nil
}

func (s *Server) Checkout(ctx context.Context, req *gen.CheckoutRequest) (*gen.CheckoutResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.Checkout(ctx, repos, req.Branch)
	return &gen.CheckoutResponse{}, nil
}

func (s *Server) Branch(ctx context.Context, req *gen.BranchRequest) (*gen.BranchResponse, error) {
	return &gen.BranchResponse{
		Branches: s.mng.Branch(ctx),
	}, nil
}

func (s *Server) Pull(ctx context.Context, req *gen.PullRequest) (*gen.PullResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.Pull(ctx, repos)
	return &gen.PullResponse{}, nil
}

func (s *Server) PreBuild(ctx context.Context, req *gen.PreBuildRequest) (*gen.PreBuildResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.PreBuild(ctx, repos)
	return &gen.PreBuildResponse{}, nil
}

func (s *Server) Build(ctx context.Context, req *gen.BuildRequest) (*gen.BuildResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.Build(ctx, repos)
	return &gen.BuildResponse{}, nil
}

func (s *Server) Run(ctx context.Context, req *gen.RunRequest) (*gen.RunResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.Run(ctx, repos)
	return &gen.RunResponse{}, nil
}

func (s *Server) Debug(ctx context.Context, req *gen.DebugRequest) (*gen.DebugResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.Debug(ctx, repos)
	return &gen.DebugResponse{}, nil
}

func (s *Server) Stop(ctx context.Context, req *gen.StopRequest) (*gen.StopResponse, error) {
	repos, err := s.getReposList(req.Repos)
	if err != nil {
		return nil, err
	}

	s.mng.Stop(ctx, repos)
	return &gen.StopResponse{}, nil
}

func (s *Server) getReposList(repos *gen.Repos) ([]string, error) {
	if repos.Group != nil {
		return s.mng.GetByGroup(*repos.Group)
	} else if len(repos.Repos) != 0 {
		return s.mng.GetByList(repos.Repos)
	} else {
		return s.mng.GetByGroup(entities.AllRepos)
	}
}
