package cmd

import (
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./clive.yml"

var defaultSettings = &settings{
	LoginCommand: []string{"bash", "--login"},
	FontSize:     22,
}

type configYaml struct {
	Settings *settingsYaml `yaml:"settings"`
	Actions  []interface{} `yaml:"actions"`
}

type config struct {
	Settings *settings
	Actions  []action
}

type settingsYaml struct {
	LoginCommand *[]string `yaml:"loginCommand"`
	FontSize     *int      `yaml:"fontSize"`
	FontFamily   *string   `yaml:"fontFamily"`
}

type settings struct {
	LoginCommand []string
	FontSize     int
	FontFamily   *string
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

	stgs := defaultSettings
	if y.Settings.LoginCommand != nil {
		stgs.LoginCommand = *y.Settings.LoginCommand
	}
	if y.Settings.FontSize != nil {
		stgs.FontSize = *y.Settings.FontSize
	}
	if y.Settings.FontFamily != nil {
		stgs.FontFamily = y.Settings.FontFamily
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
		Settings: stgs,
		Actions:  actions,
	}, nil
}
