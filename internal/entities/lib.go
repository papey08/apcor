package entities

import (
	"context"
)

type Lib struct {
	*Repo

	build Commands
}

func NewLib(
	repoUrl string,
	path string,
	build []string,
) (*Lib, error) {
	repo, err := NewRepo(repoUrl, path)
	if err != nil {
		return nil, err
	}

	return &Lib{
		Repo:  repo,
		build: NewCommands(build...),
	}, nil
}

func (l *Lib) Build(ctx context.Context) error {
	for _, cmd := range l.build {
		cmd.Exec(ctx, l.path)
	}
	return nil
}
