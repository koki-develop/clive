package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{0, 0},
			0,
		},
		{
			args{1, 0},
			1,
		},
		{
			args{0, 1},
			1,
		},
		{
			args{-1, 0},
			0,
		},
		{
			args{0, -1},
			0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := Max(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}
