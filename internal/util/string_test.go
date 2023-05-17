package util

import (
	"fmt"
	"testing"

	"github.com/charmbracelet/lipgloss"
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

func TestBorder(t *testing.T) {
	style := lipgloss.NewStyle()

	tests := []struct {
		str  string
		want string
	}{
		{
			"Hello World",
			`┌─────────────┐
│ Hello World │
└─────────────┘`,
		},
		{
			"Hello World\nGoodnight World",
			`┌─────────────────┐
│ Hello World     │
│ Goodnight World │
└─────────────────┘`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := Border(tt.str, style)
			assert.Equal(t, tt.want, got)
		})
	}
}
