package cmd

import (
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./clive.yml"

type configYaml struct {
	Actions []interface{} `yaml:"actions"`
}

type config struct {
	Actions []action
}

func loadConfig(p string) (*config, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var y configYaml
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, err
	}

	var actions []action
	for _, a := range y.Actions {
		action, err := parseAction(a)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	return &config{
		Actions: actions,
	}, nil
}
