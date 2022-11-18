package config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/koki-develop/clive/pkg/util"
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
					BrowserBin:          nil,
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
					BrowserBin:          util.String("BROWSER_BIN"),
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
`,
			&Config{
				Settings: &Settings{
					LoginCommand:        []string{"bash", "--login"},
					FontSize:            22,
					FontFamily:          nil,
					DefaultSpeed:        10,
					SkipPauseBeforeQuit: false,
					BrowserBin:          nil,
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
		input   map[string]interface{}
		want    *Config
		wantErr bool
	}{
		{
			map[string]interface{}{
				"actions": []interface{}{"pause"},
			},
			&Config{
				Settings: &Settings{
					LoginCommand:        []string{"bash", "--login"},
					FontSize:            22,
					FontFamily:          nil,
					DefaultSpeed:        10,
					SkipPauseBeforeQuit: false,
					BrowserBin:          nil,
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
			map[string]interface{}{
				"settings": map[string]interface{}{
					"loginCommand":        []interface{}{"hoge", "fuga"},
					"fontSize":            100,
					"fontFamily":          "FONT_FAMILY",
					"defaultSpeed":        200,
					"skipPauseBeforeQuit": true,
					"browserBin":          "BROWSER_BIN",
					"width":               1600,
					"height":              800,
				},
				"actions": []interface{}{"pause"},
			},
			&Config{
				Settings: &Settings{
					LoginCommand:        []string{"hoge", "fuga"},
					FontSize:            100,
					FontFamily:          util.String("FONT_FAMILY"),
					DefaultSpeed:        200,
					SkipPauseBeforeQuit: true,
					BrowserBin:          util.String("BROWSER_BIN"),
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
			map[string]interface{}{
				"actions": []interface{}{
					map[string]interface{}{"type": "Hello"},
					map[string]interface{}{"type": "Hello", "count": 10, "speed": 500},
					map[string]interface{}{"key": "enter"},
					map[string]interface{}{"key": "enter", "count": 10, "speed": 500},
					map[string]interface{}{"sleep": 1000},
					"pause",
					map[string]interface{}{"pause": nil},
					map[string]interface{}{"ctrl": "c"},
					map[string]interface{}{"ctrl": "c", "count": 10, "speed": 500},
				},
			},
			&Config{
				Settings: &Settings{
					LoginCommand: []string{"bash", "--login"},
					FontSize:     22,
					FontFamily:   nil,
					DefaultSpeed: 10,
					BrowserBin:   nil,
					Width:        nil,
					Height:       nil,
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
				},
			},
			false,
		},
		{
			map[string]interface{}{"a": "A", "actions": []interface{}{"pause"}},
			nil,
			true,
		},
		{
			map[string]interface{}{"settings": "hello world"},
			nil,
			true,
		},
		{
			map[string]interface{}{"actions": "hello world"},
			nil,
			true,
		},
		{
			map[string]interface{}{"actions": []interface{}{map[string]interface{}{"type": "hello world", "unknownField": "value"}}},
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
