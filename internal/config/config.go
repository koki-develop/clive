package config

import (
	"io"
	"os"

	"github.com/koki-develop/clive/internal/util"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type configYaml struct {
	Settings map[string]any `yaml:"settings"`
	Actions  []any          `yaml:"actions"`
}

type Config struct {
	Settings *Settings
	Actions  []Action
}

func Load(name string) (*Config, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg, err := Decode(f)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func Decode(r io.Reader) (*Config, error) {
	m := map[string]any{}
	if err := yaml.NewDecoder(r).Decode(&m); err != nil {
		return nil, err
	}

	cfg, err := DecodeMap(m)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func DecodeMap(m map[string]any) (*Config, error) {
	invalidFields := []string{}
	for k := range m {
		if !util.Contains([]string{"settings", "actions"}, k) {
			invalidFields = append(invalidFields, k)
		}
	}
	if len(invalidFields) > 0 {
		return nil, errors.WithMessagef(ErrInvalidConfig, "unknown fields %s", invalidFields)
	}

	var y configYaml
	if err := mapstructure.Decode(m, &y); err != nil {
		return nil, errors.WithMessage(ErrInvalidConfig, err.Error())
	}

	cfg, err := yamlToConfig(&y)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func yamlToConfig(y *configYaml) (*Config, error) {
	stgs, err := DecodeSettings(y.Settings)
	if err != nil {
		return nil, err
	}

	var actions []Action
	for _, a := range y.Actions {
		action, err := ParseAction(stgs, a)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	return &Config{
		Settings: stgs,
		Actions:  actions,
	}, nil
}
