package config

import (
	"fmt"
	"testing"

	"github.com/koki-develop/clive/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSettings(t *testing.T) {
	tests := []struct {
		input   map[string]interface{}
		want    *Settings
		wantErr bool
	}{
		{
			nil,
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"loginCommand": []string{"zsh", "--login"},
			},
			&Settings{
				LoginCommand:        []string{"zsh", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"fontSize": 100,
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            100,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"fontFamily": "FONT_FAMILY",
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          util.String("FONT_FAMILY"),
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"defaultSpeed": 200,
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        200,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"skipPauseBeforeQuit": true,
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: true,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"browserBin": "BROWSER_BIN",
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          util.String("BROWSER_BIN"),
				Headless:            false,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"headless": true,
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            true,
				Width:               nil,
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"width": 2000,
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               util.Int(2000),
				Height:              nil,
			},
			false,
		},
		{
			map[string]interface{}{
				"height": 1000,
			},
			&Settings{
				LoginCommand:        []string{"bash", "--login"},
				FontSize:            22,
				FontFamily:          nil,
				DefaultSpeed:        10,
				SkipPauseBeforeQuit: false,
				BrowserBin:          nil,
				Headless:            false,
				Width:               nil,
				Height:              util.Int(1000),
			},
			false,
		},
		{
			map[string]interface{}{
				"loginCommand":        []string{"zsh", "--login"},
				"fontSize":            100,
				"fontFamily":          "FONT_FAMILY",
				"defaultSpeed":        200,
				"skipPauseBeforeQuit": true,
				"browserBin":          "BROWSER_BIN",
				"headless":            true,
				"width":               2000,
				"height":              1000,
			},
			&Settings{
				LoginCommand:        []string{"zsh", "--login"},
				FontSize:            100,
				FontFamily:          util.String("FONT_FAMILY"),
				DefaultSpeed:        200,
				SkipPauseBeforeQuit: true,
				BrowserBin:          util.String("BROWSER_BIN"),
				Headless:            true,
				Width:               util.Int(2000),
				Height:              util.Int(1000),
			},
			false,
		},
		{
			map[string]interface{}{
				"a": "A",
			},
			nil,
			true,
		},
		{
			map[string]interface{}{
				"loginCommand":        []string{"zsh", "--login"},
				"fontSize":            100,
				"fontFamily":          "FONT_FAMILY",
				"defaultSpeed":        200,
				"skipPauseBeforeQuit": true,
				"browserBin":          "BROWSER_BIN",
				"headless":            true,
				"width":               2000,
				"height":              1000,
				"a":                   "A",
			},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := DecodeSettings(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
