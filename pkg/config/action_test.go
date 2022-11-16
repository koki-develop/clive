package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeAction_String(t *testing.T) {
	type fields struct {
		Type  string
		Count int
		Speed int
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields{"Hello World", 1, 10},
			"Type: Hello World",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			action := &TypeAction{
				Type:  tt.fields.Type,
				Count: tt.fields.Count,
				Speed: tt.fields.Speed,
			}
			got := action.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestKeyAction_String(t *testing.T) {
	type fields struct {
		Key   string
		Count int
		Speed int
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields{Key: "enter", Count: 1, Speed: 10},
			"Key: enter",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			action := &KeyAction{
				Key:   tt.fields.Key,
				Count: tt.fields.Count,
				Speed: tt.fields.Speed,
			}
			got := action.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSleepAction_String(t *testing.T) {
	type fields struct {
		Sleep int
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields{Sleep: 1000},
			"Sleep: 1000ms",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			action := &SleepAction{
				Sleep: tt.fields.Sleep,
			}
			got := action.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPauseAction_String(t *testing.T) {
	type fields struct{}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields{},
			"Pause: Press enter to continue",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			action := &PauseAction{}
			got := action.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCtrlAction_String(t *testing.T) {
	type fields struct {
		Ctrl  string
		Count int
		Speed int
	}
	tests := []struct {
		fields fields
		want   string
	}{
		{
			fields{Ctrl: "c", Count: 1, Speed: 10},
			"Ctrl+c",
		},
		{
			fields{Ctrl: "cccc", Count: 1, Speed: 10},
			"Ctrl+cccc",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			action := &CtrlAction{
				Ctrl:  tt.fields.Ctrl,
				Count: tt.fields.Count,
				Speed: tt.fields.Speed,
			}
			got := action.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseTypeAction(t *testing.T) {
	stgs := &Settings{DefaultSpeed: 10}

	tests := []struct {
		input   map[string]interface{}
		want    Action
		wantErr bool
	}{
		{
			map[string]interface{}{"type": "Hello World"},
			&TypeAction{Type: "Hello World", Count: 1, Speed: stgs.DefaultSpeed},
			false,
		},
		{
			map[string]interface{}{"type": "Hello World", "count": 10},
			&TypeAction{Type: "Hello World", Count: 10, Speed: stgs.DefaultSpeed},
			false,
		},
		{
			map[string]interface{}{"type": "Hello World", "speed": 500},
			&TypeAction{Type: "Hello World", Count: 1, Speed: 500},
			false,
		},
		{
			map[string]interface{}{"type": "Hello World", "count": 10, "speed": 500},
			&TypeAction{Type: "Hello World", Count: 10, Speed: 500},
			false,
		},
		{
			map[string]interface{}{"type": "Hello World", "a": "A"},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseTypeAction(stgs, tt.input)
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
	stgs := &Settings{DefaultSpeed: 10}

	tests := []struct {
		input   map[string]interface{}
		want    Action
		wantErr bool
	}{
		{
			map[string]interface{}{"key": "enter"},
			&KeyAction{Key: "enter", Count: 1, Speed: stgs.DefaultSpeed},
			false,
		},
		{
			map[string]interface{}{"key": "enter", "count": 10},
			&KeyAction{Key: "enter", Count: 10, Speed: stgs.DefaultSpeed},
			false,
		},
		{
			map[string]interface{}{"key": "enter", "speed": 500},
			&KeyAction{Key: "enter", Count: 1, Speed: 500},
			false,
		},
		{
			map[string]interface{}{"key": "enter", "count": 10, "speed": 500},
			&KeyAction{Key: "enter", Count: 10, Speed: 500},
			false,
		},
		{
			map[string]interface{}{"key": "enter", "a": "A"},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseKeyAction(stgs, tt.input)
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
	tests := []struct {
		input   map[string]interface{}
		want    Action
		wantErr bool
	}{
		{
			map[string]interface{}{"sleep": 3000},
			&SleepAction{Sleep: 3000},
			false,
		},
		{
			map[string]interface{}{"sleep": 3000, "a": "A"},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseSleepAction(nil, tt.input)
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
	tests := []struct {
		input   map[string]interface{}
		want    Action
		wantErr bool
	}{
		{
			map[string]interface{}{"pause": nil},
			&PauseAction{},
			false,
		},
		{
			map[string]interface{}{"pause": struct{}{}},
			&PauseAction{},
			false,
		},
		{
			map[string]interface{}{"pause": nil, "a": "A"},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parsePauseAction(nil, tt.input)
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
	stgs := &Settings{DefaultSpeed: 10}

	tests := []struct {
		input   map[string]interface{}
		want    Action
		wantErr bool
	}{
		{
			map[string]interface{}{"ctrl": "c"},
			&CtrlAction{Ctrl: "c", Count: 1, Speed: 10},
			false,
		},
		{
			map[string]interface{}{"ctrl": "c", "count": 10},
			&CtrlAction{Ctrl: "c", Count: 10, Speed: stgs.DefaultSpeed},
			false,
		},
		{
			map[string]interface{}{"ctrl": "c", "speed": 500},
			&CtrlAction{Ctrl: "c", Count: 1, Speed: 500},
			false,
		},
		{
			map[string]interface{}{"ctrl": "c", "count": 10, "speed": 500},
			&CtrlAction{Ctrl: "c", Count: 10, Speed: 500},
			false,
		},
		{
			map[string]interface{}{"ctrl": "c", "a": "A"},
			nil,
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseCtrlAction(stgs, tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
