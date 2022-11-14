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
		args  args
		want  string
		want1 bool
	}{
		{
			args{"aaaaaaaaa", 10},
			"aaaaaaaaa",
			false,
		},
		{
			args{"aaaaaaaaaa", 10},
			"aaaaaaaaaa",
			false,
		},
		{
			args{"aaaaaaaaaaa", 10},
			"aaaaaaaaaa",
			true,
		},
		{
			args{"aaaaaaaaa\naaaaa", 10},
			"aaaaaaaaa",
			true,
		},
		{
			args{"aaaaaaaaaa\naaaaa", 10},
			"aaaaaaaaaa",
			true,
		},
		{
			args{"aaaaaaaaaaa\naaaaa", 10},
			"aaaaaaaaaa",
			true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, got1 := TruncateString(tt.args.s, tt.args.l)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestPaddingRight(t *testing.T) {
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
			got := PaddingRight(tt.args.s, tt.args.l)
			assert.Equal(t, tt.want, got)
		})
	}

}

func TestString(t *testing.T) {
	s := "hello"
	assert.Equal(t, &s, String(s))
}

func TestContains(t *testing.T) {
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
			got := Contains(tt.args.slice, tt.args.r)
			assert.Equal(t, tt.want, got)
		})
	}

}
