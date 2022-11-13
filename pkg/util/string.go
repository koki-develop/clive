package util

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

func TruncateString(s string, l int) string {
	rows := strings.Split(s, "\n")
	trunc := false
	if len(rows) > 1 {
		s = rows[0]
		trunc = true
	}
	if utf8.RuneCountInString(s) > l {
		s = string([]rune(s)[:l])
		trunc = true
	}
	if trunc {
		s += lipgloss.NewStyle().Faint(true).Render("...")
	}

	return s
}
