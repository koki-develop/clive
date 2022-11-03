package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTypeAction(t *testing.T) {
	type args struct {
		settings *settings
		m        map[string]interface{}
	}
	tests := []struct {
		args    args
		want    *typeAction
		wantErr bool
	}{
		{
			args{
				&settings{DefaultSpeed: 10},
				map[string]interface{}{
					"type": "Hello World",
				},
			},
			&typeAction{
				Type:  "Hello World",
				Count: 1,
				Speed: 10,
			},
			false,
		},
		{
			args{
				&settings{DefaultSpeed: 10},
				map[string]interface{}{
					"type":  "Hello World",
					"count": 10,
					"speed": 500,
				},
			},
			&typeAction{
				Type:  "Hello World",
				Count: 10,
				Speed: 500,
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseTypeAction(tt.args.settings, tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseKeyAction(t *testing.T) {
	type args struct {
		settings *settings
		m        map[string]interface{}
	}
	tests := []struct {
		args    args
		want    *keyAction
		wantErr bool
	}{
		{
			args{
				&settings{DefaultSpeed: 10},
				map[string]interface{}{
					"key": "enter",
				},
			},
			&keyAction{
				Key:   "enter",
				Count: 1,
				Speed: 10,
			},
			false,
		},
		{
			args{
				&settings{DefaultSpeed: 10},
				map[string]interface{}{
					"key":   "enter",
					"count": 10,
					"speed": 500,
				},
			},
			&keyAction{
				Key:   "enter",
				Count: 10,
				Speed: 500,
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseKeyAction(tt.args.settings, tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseSleepAction(t *testing.T) {
	type args struct {
		settings *settings
		m        map[string]interface{}
	}
	tests := []struct {
		args    args
		want    *sleepAction
		wantErr bool
	}{
		{
			args{
				nil,
				map[string]interface{}{
					"sleep": 3000,
				},
			},
			&sleepAction{
				Sleep: 3000,
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseSleepAction(tt.args.settings, tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parsePauseAction(t *testing.T) {
	type args struct {
		settings *settings
		m        map[string]interface{}
	}
	tests := []struct {
		args    args
		want    *pauseAction
		wantErr bool
	}{
		{
			args{
				nil,
				map[string]interface{}{
					"pause": nil,
				},
			},
			&pauseAction{},
			false,
		},
		{
			args{
				nil,
				map[string]interface{}{
					"pause": struct{}{},
				},
			},
			&pauseAction{},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parsePauseAction(tt.args.settings, tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseCtrlAction(t *testing.T) {
	type args struct {
		settings *settings
		m        map[string]interface{}
	}
	tests := []struct {
		args    args
		want    *ctrlAction
		wantErr bool
	}{
		{
			args{
				&settings{DefaultSpeed: 10},
				map[string]interface{}{
					"ctrl": "c",
				},
			},
			&ctrlAction{
				Ctrl:  "c",
				Count: 1,
				Speed: 10,
			},
			false,
		},
		{
			args{
				&settings{DefaultSpeed: 10},
				map[string]interface{}{
					"ctrl":  "c",
					"count": 10,
					"speed": 500,
				},
			},
			&ctrlAction{
				Ctrl:  "c",
				Count: 10,
				Speed: 500,
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseCtrlAction(tt.args.settings, tt.args.m)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
