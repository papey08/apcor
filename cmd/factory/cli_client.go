package factory

import "github.com/papey08/apcor/internal/cli"

func CliClient(addr string) (*cli.Client, error) {
	return cli.NewClient(addr)
}
