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

func TestDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{0},
			1,
		},
		{
			args{1},
			1,
		},
		{
			args{10},
			2,
		},
		{
			args{100},
			3,
		},
		{
			args{1000},
			4,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := Digits(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt(t *testing.T) {
	i := 1
	assert.Equal(t, &i, Int(i))
}
