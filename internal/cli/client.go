package cli

import (
	"context"

	"github.com/papey08/apcor/internal/cli/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client gen.CliServiceClient
}

func NewClient(addr string) (*Client, error) {
	if conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		return &Client{
			client: gen.NewCliServiceClient(conn),
		}, nil
	}
}

func (c *Client) InstallDeps(ctx context.Context, platform string) error {
	if _, err := c.client.InstallDeps(ctx, &gen.InstallDepsRequest{
		Platform: platform,
	}); err != nil {
		return err
	}

	return nil
}

func (c *Client) ExecPreTasks(ctx context.Context) error {
	if _, err := c.client.ExecPreTasks(ctx, &gen.ExecPreTasksRequest{}); err != nil {
		return err
	}

	return nil
}

func (c *Client) ExecPostTasks(ctx context.Context) error {
	if _, err := c.client.ExecPostTasks(ctx, &gen.ExecPostTasksRequest{}); err != nil {
		return err
	}

	return nil
}

func (c *Client) Clone(ctx context.Context, repos []string, group *string) error {
	req := c.createReposList(repos, group)
	if _, err := c.client.Clone(ctx, &gen.CloneRequest{
		Repos: req,
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) Checkout(ctx context.Context, repos []string, group *string, branch string) error {
	req := c.createReposList(repos, group)
	_, err := c.client.Checkout(ctx, &gen.CheckoutRequest{
		Repos:  req,
		Branch: branch,
	})
	return err
}

func (c *Client) Branch(ctx context.Context) (map[string]string, error) {
	if resp, err := c.client.Branch(ctx, &gen.BranchRequest{}); err != nil {
		return nil, err
	} else {
		return resp.Branches, nil
	}
}

func (c *Client) Pull(ctx context.Context, repos []string, group *string) error {
	req := c.createReposList(repos, group)
	if _, err := c.client.Pull(ctx, &gen.PullRequest{
		Repos: req,
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) PreBuild(ctx context.Context, repos []string, group *string) error {
	req := c.createReposList(repos, group)
	if _, err := c.client.PreBuild(ctx, &gen.PreBuildRequest{
		Repos: req,
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) Build(ctx context.Context, repos []string, group *string) error {
	req := c.createReposList(repos, group)
	if _, err := c.client.Build(ctx, &gen.BuildRequest{
		Repos: req,
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) Run(ctx context.Context, repos []string, group *string) error {
	req := c.createReposList(repos, group)
	if _, err := c.client.Run(ctx, &gen.RunRequest{
		Repos: req,
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) Debug(ctx context.Context, repos []string, group *string) error {
	req := c.createReposList(repos, group)
	if _, err := c.client.Debug(ctx, &gen.DebugRequest{
		Repos: req,
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) Stop(ctx context.Context, repos []string, group *string) error {
	req := c.createReposList(repos, group)
	if _, err := c.client.Stop(ctx, &gen.StopRequest{
		Repos: req,
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) createReposList(repos []string, group *string) *gen.Repos {
	if group != nil {
		return &gen.Repos{
			Group: group,
		}
	} else {
		return &gen.Repos{
			Repos: repos,
		}
	}
}
