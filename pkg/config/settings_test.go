package config

import (
	"fmt"
	"testing"

	"github.com/koki-develop/clive/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSettings(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		args    args
		want    *Settings
		wantErr bool
	}{
		{
			args{nil},
			&Settings{
				LoginCommand: []string{"bash", "--login"},
				FontSize:     22,
				FontFamily:   nil,
				DefaultSpeed: 10,
				BrowserBin:   nil,
			},
			false,
		},
		{
			args{map[string]interface{}{}},
			&Settings{
				LoginCommand: []string{"bash", "--login"},
				FontSize:     22,
				FontFamily:   nil,
				DefaultSpeed: 10,
				BrowserBin:   nil,
			},
			false,
		},
		{
			args{map[string]interface{}{
				"loginCommand": []string{"zsh", "--login"},
			}},
			&Settings{
				LoginCommand: []string{"zsh", "--login"},
				FontSize:     22,
				FontFamily:   nil,
				DefaultSpeed: 10,
				BrowserBin:   nil,
			},
			false,
		},
		{
			args{map[string]interface{}{
				"loginCommand": []string{"zsh", "--login"},
				"fontSize":     100,
				"fontFamily":   "FONT_FAMILY",
				"defaultSpeed": 200,
				"browserBin":   "BROWSER_BIN",
			}},
			&Settings{
				LoginCommand: []string{"zsh", "--login"},
				FontSize:     100,
				FontFamily:   util.String("FONT_FAMILY"),
				DefaultSpeed: 200,
				BrowserBin:   util.String("BROWSER_BIN"),
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := DecodeSettings(tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
