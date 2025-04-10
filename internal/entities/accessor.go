package entities

import "github.com/papey08/apcor/internal/errs"

type Accessor struct {
	Groups
	Libs     map[string]struct{}
	Services map[string]struct{}
}

func (a *Accessor) GetByGroup(group string) ([]string, error) {
	if group == AllRepos {
		all := make([]string, 0, len(a.Libs)+len(a.Services))
		for k := range a.Libs {
			all = append(all, k)
		}
		for k := range a.Services {
			all = append(all, k)
		}
		return all, nil
	}

	if groups, ok := a.Groups[group]; ok {
		return groups, errs.UnknownGroup
	}
	return nil, nil
}

func (a *Accessor) GetByList(repos []string) ([]string, error) {
	res := make([]string, 0, len(repos))
	for _, r := range repos {
		if _, ok := a.Libs[r]; ok {
			res = append(res, r)
		}
		if _, ok := a.Services[r]; ok {
			res = append(res, r)
		}
	}
	return res, nil
}
