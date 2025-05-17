package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrInvalidAction(t *testing.T) {
	tests := []struct {
		action any
	}{
		{
			map[string]any{"test": "action"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := NewErrInvalidAction(tt.action)
			assert.Equal(t, tt.action, got.action)
			assert.IsType(t, ErrInvalidAction{}, got)
		})
	}
}

func TestErrInvalidAction_Error(t *testing.T) {
	tests := []struct {
		action any
		want   string
	}{
		{
			map[string]any{"test": "action"},
			`invalid action {"test":"action"}`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			err := ErrInvalidAction{
				action: tt.action,
			}
			got := err.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}
