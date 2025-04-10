package entities

import (
	"context"
	"errors"
	"net/url"
	"os/exec"

	"github.com/papey08/apcor/internal/errs"
)

type Repo struct {
	remoteRepo *url.URL
	path       string
}

func NewRepo(repoUrl string, path string) (*Repo, error) {
	u, err := url.Parse(repoUrl)
	if err != nil {
		return nil, errors.Join(errs.InvalidUrl, err)
	}
	return &Repo{
		remoteRepo: u,
		path:       path,
	}, nil
}

func (r *Repo) Clone(ctx context.Context) error {
	return exec.CommandContext(ctx, "bash", scriptPathClone, r.remoteRepo.String(), r.path).Run()
}

func (r *Repo) Checkout(ctx context.Context, branch string) error {
	cmd := exec.CommandContext(ctx, "bash", scriptPathCheckout, branch)
	cmd.Dir = r.path
	return cmd.Run()
}

func (r *Repo) Branch(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "bash", scriptPathBranch)
	cmd.Dir = r.path
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (r *Repo) Pull(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "bash", scriptPathPull)
	cmd.Dir = r.path
	return cmd.Run()
}

const (
	scriptPathClone    = "./scripts/clone.sh"
	scriptPathCheckout = "./scripts/checkout.sh"
	scriptPathBranch   = "./scripts/branch.sh"
	scriptPathPull     = "./scripts/pull.sh"
)
