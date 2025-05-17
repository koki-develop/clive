package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/koki-develop/clive/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_Decode(t *testing.T) {
	tests := []struct {
		yaml    string
		want    *Config
		wantErr bool
	}{
		{
			`
actions:
  - pause
`,
			&Config{
				Settings: &Settings{
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
				},
				Actions: []Action{
					&PauseAction{},
				},
			},
			false,
		},

		{
			`
settings:
  loginCommand: ["hoge", "fuga"]
  fontSize: 100
  fontFamily: FONT_FAMILY
  defaultSpeed: 200
  skipPauseBeforeQuit: true
  screenshotsDir: SCREENSHOTS_DIR
  headless: true
  browserBin: BROWSER_BIN
  width: 1600
  height: 800
actions:
  - pause
`,
			&Config{
				Settings: &Settings{
					LoginCommand:        []string{"hoge", "fuga"},
					FontSize:            100,
					FontFamily:          util.String("FONT_FAMILY"),
					DefaultSpeed:        200,
					SkipPauseBeforeQuit: true,
					ScreenshotsDir:      "SCREENSHOTS_DIR",
					BrowserBin:          util.String("BROWSER_BIN"),
					Headless:            true,
					Width:               util.Int(1600),
					Height:              util.Int(800),
				},
				Actions: []Action{
					&PauseAction{},
				},
			},
			false,
		},
		{
			`
actions:
  - type: Hello
  - type: Hello
    count: 10
    speed: 500
  - key: enter
  - key: enter
    count: 10
    speed: 500
  - sleep: 1000
  - pause
  - pause:
  - ctrl: c
  - ctrl: c
    count: 10
    speed: 500
  - screenshot
  - screenshot:
`,
			&Config{
				Settings: &Settings{
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
				},
				Actions: []Action{
					&TypeAction{Type: "Hello", Count: 1, Speed: 10},
					&TypeAction{Type: "Hello", Count: 10, Speed: 500},
					&KeyAction{Key: "enter", Count: 1, Speed: 10},
					&KeyAction{Key: "enter", Count: 10, Speed: 500},
					&SleepAction{Sleep: 1000},
					&PauseAction{},
					&PauseAction{},
					&CtrlAction{Ctrl: "c", Count: 1, Speed: 10},
					&CtrlAction{Ctrl: "c", Count: 10, Speed: 500},
					&ScreenshotAction{},
					&ScreenshotAction{},
				},
			},
			false,
		},
		{
			`
a: A
actions:
  - pause
`,
			nil,
			true,
		},
		{
			"settings: hello world",
			nil,
			true,
		},
		{
			"actions: hello world",
			nil,
			true,
		},
		{
			`
actions:
  - type: hello world
    unknownField: value
`,
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := Decode(strings.NewReader(tt.yaml))

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDecodeMap(t *testing.T) {
	tests := []struct {
		input   map[string]any
		want    *Config
		wantErr bool
	}{
		{
			map[string]any{
				"actions": []any{"pause"},
			},
			&Config{
				Settings: &Settings{
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
				},
				Actions: []Action{
					&PauseAction{},
				},
			},
			false,
		},
		{
			map[string]any{
				"settings": map[string]any{
					"loginCommand":        []any{"hoge", "fuga"},
					"fontSize":            100,
					"fontFamily":          "FONT_FAMILY",
					"defaultSpeed":        200,
					"skipPauseBeforeQuit": true,
					"screenshotsDir":      "SCREENSHOTS_DIR",
					"browserBin":          "BROWSER_BIN",
					"headless":            true,
					"width":               1600,
					"height":              800,
				},
				"actions": []any{"pause"},
			},
			&Config{
				Settings: &Settings{
					LoginCommand:        []string{"hoge", "fuga"},
					FontSize:            100,
					FontFamily:          util.String("FONT_FAMILY"),
					DefaultSpeed:        200,
					SkipPauseBeforeQuit: true,
					ScreenshotsDir:      "SCREENSHOTS_DIR",
					BrowserBin:          util.String("BROWSER_BIN"),
					Headless:            true,
					Width:               util.Int(1600),
					Height:              util.Int(800),
				},
				Actions: []Action{
					&PauseAction{},
				},
			},
			false,
		},
		{
			map[string]any{
				"actions": []any{
					map[string]any{"type": "Hello"},
					map[string]any{"type": "Hello", "count": 10, "speed": 500},
					map[string]any{"key": "enter"},
					map[string]any{"key": "enter", "count": 10, "speed": 500},
					map[string]any{"sleep": 1000},
					"pause",
					map[string]any{"pause": nil},
					map[string]any{"pause": struct{}{}},
					map[string]any{"ctrl": "c"},
					map[string]any{"ctrl": "c", "count": 10, "speed": 500},
					"screenshot",
					map[string]any{"screenshot": nil},
					map[string]any{"screenshot": "SCREENSHOT"},
				},
			},
			&Config{
				Settings: &Settings{
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
				},
				Actions: []Action{
					&TypeAction{Type: "Hello", Count: 1, Speed: 10},
					&TypeAction{Type: "Hello", Count: 10, Speed: 500},
					&KeyAction{Key: "enter", Count: 1, Speed: 10},
					&KeyAction{Key: "enter", Count: 10, Speed: 500},
					&SleepAction{Sleep: 1000},
					&PauseAction{},
					&PauseAction{},
					&PauseAction{},
					&CtrlAction{Ctrl: "c", Count: 1, Speed: 10},
					&CtrlAction{Ctrl: "c", Count: 10, Speed: 500},
					&ScreenshotAction{},
					&ScreenshotAction{},
					&ScreenshotAction{Screenshot: util.String("SCREENSHOT")},
				},
			},
			false,
		},
		{
			map[string]any{"a": "A", "actions": []any{"pause"}},
			nil,
			true,
		},
		{
			map[string]any{"settings": "hello world"},
			nil,
			true,
		},
		{
			map[string]any{"actions": "hello world"},
			nil,
			true,
		},
		{
			map[string]any{"actions": []any{map[string]any{"type": "hello world", "unknownField": "value"}}},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := DecodeMap(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
