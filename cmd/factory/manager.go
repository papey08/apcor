package factory

import (
	"github.com/papey08/apcor/internal/parser"

	"github.com/papey08/apcor/internal/manager"
)

func Manager(path string) (*manager.Manager, error) {
	data, err := parser.ParseData(path)
	if err != nil {
		return nil, err
	}
	return manager.NewManager(data)
}
