package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrInvalidAction(t *testing.T) {
	tests := []struct {
		action interface{}
	}{
		{
			map[string]interface{}{"test": "action"},
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
