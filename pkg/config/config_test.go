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
					LoginCommand: []string{"bash", "--login"},
					FontSize:     22,
					FontFamily:   nil,
					DefaultSpeed: 10,
					BrowserBin:   nil,
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
  fontFamily: FontName
  defaultSpeed: 200
  browserBin: /path/to/browser
actions:
  - pause
`,
			&Config{
				Settings: &Settings{
					LoginCommand: []string{"hoge", "fuga"},
					FontSize:     100,
					FontFamily:   util.String("FontName"),
					DefaultSpeed: 200,
					BrowserBin:   util.String("/path/to/browser"),
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
					LoginCommand: []string{"bash", "--login"},
					FontSize:     22,
					FontFamily:   nil,
					DefaultSpeed: 10,
					BrowserBin:   nil,
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
