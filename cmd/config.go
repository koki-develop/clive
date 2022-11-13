package cmd

import (
	"io"
	"os"

	"github.com/koki-develop/clive/pkg/config"
	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./clive.yml"

type configYaml struct {
	Settings map[string]interface{} `yaml:"settings"`
	Actions  []interface{}          `yaml:"actions"`
}

type legacyConfig struct {
	Settings *config.Settings
	Actions  []action
}

func loadConfig(p string) (*legacyConfig, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg, err := decodeConfig(f)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func decodeConfig(f io.Reader) (*legacyConfig, error) {
	var y configYaml
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, err
	}

	settings, err := config.DecodeSettings(y.Settings)
	if err != nil {
		return nil, err
	}

	var actions []action
	for _, a := range y.Actions {
		action, err := parseAction(settings, a)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	return &legacyConfig{
		Settings: settings,
		Actions:  actions,
	}, nil
}
