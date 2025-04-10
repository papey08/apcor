package factory

import (
	"github.com/papey08/apcor/internal/cli"
	"google.golang.org/grpc"
)

func CliServer(path string) (*grpc.Server, error) {
	manager, err := Manager(path)
	if err != nil {
		return nil, err
	}

	return cli.NewServer(manager), nil
}
