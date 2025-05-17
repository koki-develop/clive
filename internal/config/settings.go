package config

import (
	"github.com/koki-develop/clive/internal/util"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type Settings struct {
	LoginCommand        []string `mapstructure:"loginCommand"`
	FontSize            int      `mapstructure:"fontSize"`
	FontFamily          *string  `mapstructure:"fontFamily"`
	DefaultSpeed        int      `mapstructure:"defaultSpeed"`
	SkipPauseBeforeQuit bool     `mapstructure:"skipPauseBeforeQuit"`
	ScreenshotsDir      string   `mapstructure:"screenshotsDir"`
	BrowserBin          *string  `mapstructure:"browserBin"`
	Headless            bool     `mapstructure:"headless"`
	Width               *int     `mapstructure:"width"`
	Height              *int     `mapstructre:"height"`
}

var settingsFields = []string{
	"loginCommand",
	"fontSize",
	"fontFamily",
	"defaultSpeed",
	"skipPauseBeforeQuit",
	"screenshotsDir",
	"browserBin",
	"headless",
	"width",
	"height",
}

func DecodeSettings(m map[string]any) (*Settings, error) {
	stgs := Settings{
		LoginCommand:        []string{"bash", "--login"},
		FontSize:            22,
		FontFamily:          nil,
		DefaultSpeed:        10,
		SkipPauseBeforeQuit: false,
		ScreenshotsDir:      "screenshots",
		BrowserBin:          nil,
		Headless:            false,
		Width:               nil,
		Height:              nil,
	}
	if m == nil {
		return &stgs, nil
	}

	invalidFields := []string{}
	for k := range m {
		if !util.Contains(settingsFields, k) {
			invalidFields = append(invalidFields, k)
		}
	}
	if len(invalidFields) > 0 {
		return nil, errors.WithMessagef(ErrInvalidSettings, "unknown fields %s", invalidFields)
	}

	if err := mapstructure.Decode(m, &stgs); err != nil {
		return nil, errors.WithMessagef(ErrInvalidSettings, "%s", err)
	}

	return &stgs, nil
}
