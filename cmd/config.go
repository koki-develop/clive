package cmd

import (
	"io"
	"os"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./clive.yml"

type configYaml struct {
	Settings map[string]interface{} `yaml:"settings"`
	Actions  []interface{}          `yaml:"actions"`
}

type config struct {
	Settings *settings
	Actions  []action
}

type settings struct {
	LoginCommand []string `mapstructure:"loginCommand"`
	FontSize     int      `mapstructure:"fontSize"`
	FontFamily   *string  `mapstructure:"fontFamily"`
	DefaultSpeed int      `mapstructure:"defaultSpeed"`
}

func loadConfig(p string) (*config, error) {
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

func decodeConfig(f io.Reader) (*config, error) {
	var y configYaml
	if err := yaml.NewDecoder(f).Decode(&y); err != nil {
		return nil, err
	}

	settings := settings{
		LoginCommand: []string{"bash", "--login"},
		FontSize:     22,
		DefaultSpeed: 10,
	}
	if err := mapstructure.Decode(y.Settings, &settings); err != nil {
		return nil, err
	}

	var actions []action
	for _, a := range y.Actions {
		action, err := parseAction(&settings, a)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}

	return &config{
		Settings: &settings,
		Actions:  actions,
	}, nil
}
