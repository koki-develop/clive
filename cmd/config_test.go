package cmd

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeConfig(t *testing.T) {
	type args struct {
		f io.Reader
	}
	tests := []struct {
		args    args
		want    *config
		wantErr bool
	}{
		/*
		 * Settings
		 */
		{
			args{strings.NewReader(`
actions:
  - pause
`)},
			&config{
				Settings: &settings{
					LoginCommand: []string{"bash", "--login"},
					FontSize:     22,
					FontFamily:   nil,
					DefaultSpeed: 10,
					BrowserBin:   nil,
				},
				Actions: []action{
					&pauseAction{},
				},
			},
			false,
		},
		{
			args{strings.NewReader(`
settings:
  loginCommand: ["hoge", "fuga"]
  fontSize: 999
  fontFamily: FontName
  defaultSpeed: 999
  browserBin: /path/to/browser
actions:
  - pause
`)},
			&config{
				Settings: &settings{
					LoginCommand: []string{"hoge", "fuga"},
					FontSize:     999,
					FontFamily:   ptr("FontName"),
					DefaultSpeed: 999,
					BrowserBin:   ptr("/path/to/browser"),
				},
				Actions: []action{
					&pauseAction{},
				},
			},
			false,
		},
		/*
		 * Actions
		 */
		{
			args{strings.NewReader(`
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
`)},
			&config{
				Settings: &settings{
					LoginCommand: []string{"bash", "--login"},
					FontSize:     22,
					FontFamily:   nil,
					DefaultSpeed: 10,
					BrowserBin:   nil,
				},
				Actions: []action{
					&typeAction{Type: "Hello", Count: 1, Speed: 10},
					&typeAction{Type: "Hello", Count: 10, Speed: 500},
					&keyAction{Key: "enter", Count: 1, Speed: 10},
					&keyAction{Key: "enter", Count: 10, Speed: 500},
					&sleepAction{Sleep: 1000},
					&pauseAction{},
					&pauseAction{},
					&ctrlAction{Ctrl: "c", Count: 1, Speed: 10},
					&ctrlAction{Ctrl: "c", Count: 10, Speed: 500},
				},
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := decodeConfig(tt.args.f)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
