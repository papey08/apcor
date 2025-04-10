package parser

import (
	"errors"
	"os"

	"github.com/papey08/apcor/internal/errs"
	"gopkg.in/yaml.v1"
)

type (
	ApcorParsingData struct {
		ApcorVersion string `yaml:"version"`
		ProjectName  string `yaml:"name"`

		// InstallDeps is a map where key is OS name (windows/macos/etc) and value is a list of commands to install dependencies
		InstallDeps map[string][]string `yaml:"install-deps"`

		Common `yaml:"common"`

		Libs     map[string]Lib     `yaml:"libs"`
		Services map[string]Service `yaml:"services"`

		// Groups is a map where key is a group name and value is a list of service names
		Groups map[string][]string `yaml:"groups"`
	}

	Common struct {
		PreTasks  []string `yaml:"pre-tasks"`
		PostTasks []string `yaml:"post-tasks"`
	}

	Lib struct {
		RemoteRepoUrl string   `yaml:"remote-repo-url"`
		Path          string   `yaml:"path"`
		Build         []string `yaml:"build"`
	}

	Service struct {
		RemoteRepoUrl string   `yaml:"remote-repo-url"`
		Path          string   `yaml:"path"`
		PreBuild      []string `yaml:"pre-build"`
		Build         []string `yaml:"build"`
		Run           string   `yaml:"run"`
		Debug         string   `yaml:"debug"`
		Stop          string   `yaml:"stop"`
	}
)

func ParseData(path string) (*ApcorParsingData, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Join(errs.ReadingConfigData, err)
	}

	var data ApcorParsingData
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, errors.Join(errs.ReadingConfigData, err)
	}

	return &data, nil
}
