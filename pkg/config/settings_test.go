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
				BrowserBin:          nil,
				SkipPauseBeforeQuit: false,
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
				BrowserBin:          nil,
				SkipPauseBeforeQuit: false,
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
				BrowserBin:          nil,
				SkipPauseBeforeQuit: false,
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
				BrowserBin:          nil,
				SkipPauseBeforeQuit: false,
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
				BrowserBin:          nil,
				SkipPauseBeforeQuit: false,
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
				BrowserBin:          nil,
				SkipPauseBeforeQuit: false,
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
				BrowserBin:          util.String("BROWSER_BIN"),
				SkipPauseBeforeQuit: false,
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
				BrowserBin:          nil,
				SkipPauseBeforeQuit: true,
			},
			false,
		},
		{
			map[string]interface{}{
				"loginCommand":        []string{"zsh", "--login"},
				"fontSize":            100,
				"fontFamily":          "FONT_FAMILY",
				"defaultSpeed":        200,
				"browserBin":          "BROWSER_BIN",
				"skipPauseBeforeQuit": true,
			},
			&Settings{
				LoginCommand:        []string{"zsh", "--login"},
				FontSize:            100,
				FontFamily:          util.String("FONT_FAMILY"),
				DefaultSpeed:        200,
				BrowserBin:          util.String("BROWSER_BIN"),
				SkipPauseBeforeQuit: true,
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
				"browserBin":          "BROWSER_BIN",
				"skipPauseBeforeQuit": true,
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
