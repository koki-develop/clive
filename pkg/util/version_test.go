package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_Newer(t *testing.T) {
	tests := []struct {
		v    Version
		r    Version
		want bool
	}{
		{"v2.0.0", "v1.0.0", true},
		{"v1.1.0", "v1.0.0", true},
		{"v1.0.1", "v1.0.0", true},
		{"v1.0.0", "v2.0.0", false},
		{"v1.0.0", "v1.1.0", false},
		{"v1.0.0", "v1.0.1", false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := tt.v.Newer(tt.r)
			assert.Equal(t, tt.want, got)
		})
	}
}
