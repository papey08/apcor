package entities

import (
	"context"
	"os"
	"os/exec"

	"github.com/papey08/apcor/internal/errs"
	"github.com/papey08/apcor/pkg/errstream"
)

type (
	Command string

	Commands []*Command
)

func NewCommand(cmd string) *Command {
	return (*Command)(&cmd)
}

func NewCommands(cmds ...string) Commands {
	var c Commands
	for _, cmd := range cmds {
		c = append(c, NewCommand(cmd))
	}
	return c
}

func (c *Command) Exec(ctx context.Context, path string) error {
	if c == nil {
		return errs.InvalidCommand
	}
	cmd := exec.CommandContext(ctx, "sh", "-c", string(*c))
	cmd.Dir = path
	return cmd.Run()
}

func (c *Command) ExecAsync(ctx context.Context, path string, es *errstream.ErrStream) {
	if es == nil {
		go c.Exec(ctx, path)
	} else {
		es.Go(func() error {
			return c.Exec(ctx, path)
		})
	}
}

func (c *Command) ExecInteractive(ctx context.Context, path string) error {
	if c == nil {
		return errs.InvalidCommand
	}

	cmd := exec.CommandContext(ctx, "sh", "-c", string(*c))
	cmd.Dir = path
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
