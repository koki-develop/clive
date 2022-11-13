package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateString(t *testing.T) {
	type args struct {
		s string
		l int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"aaaaaaaaa", 10},
			"aaaaaaaaa",
		},
		{
			args{"aaaaaaaaaa", 10},
			"aaaaaaaaaa",
		},
		{
			args{"aaaaaaaaaaa", 10},
			"aaaaaaaaaa\x1b[2m...\x1b[0m",
		},
		{
			args{"aaaaaaaaa\naaaaa", 10},
			"aaaaaaaaa\x1b[2m...\x1b[0m",
		},
		{
			args{"aaaaaaaaaa\naaaaa", 10},
			"aaaaaaaaaa\x1b[2m...\x1b[0m",
		},
		{
			args{"aaaaaaaaaaa\naaaaa", 10},
			"aaaaaaaaaa\x1b[2m...\x1b[0m",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := TruncateString(tt.args.s, tt.args.l)
			assert.Equal(t, tt.want, got)
		})
	}
}
