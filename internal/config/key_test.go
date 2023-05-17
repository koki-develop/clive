package config

import (
	"fmt"
	"testing"

	"github.com/go-rod/rod/lib/input"
	"github.com/stretchr/testify/assert"
)

func Test_shift(t *testing.T) {
	tests := []struct {
		key  input.Key
		want input.Key
	}{
		{input.KeyA, 'A'},
		{input.Digit1, '!'},
		{input.Semicolon, ':'},
		{input.Quote, '"'},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := shift(tt.key)
			assert.Equal(t, tt.want, got)
		})
	}
}
