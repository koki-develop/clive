package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_paddingRight(t *testing.T) {
	type args struct {
		s string
		l int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"a", 4},
			"a   ",
		},
		{
			args{"aa", 4},
			"aa  ",
		},
		{
			args{"aaa", 4},
			"aaa ",
		},
		{
			args{"aaaa", 4},
			"aaaa",
		},
		{
			args{"aaaaa", 4},
			"aaaaa",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := paddingRight(tt.args.s, tt.args.l)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		slice []string
		r     string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{[]string{"aaa", "bbb", "ccc"}, "aaa"}, true},
		{args{[]string{"aaa", "bbb", "ccc"}, "bbb"}, true},
		{args{[]string{"aaa", "bbb", "ccc"}, "ccc"}, true},
		{args{[]string{"aaa", "bbb", "ccc"}, ""}, false},
		{args{[]string{"aaa", "bbb", "ccc"}, "ddd"}, false},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := contains(tt.args.slice, tt.args.r)
			assert.Equal(t, tt.want, got)
		})
	}
}
