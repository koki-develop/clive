package util

import (
	"bytes"
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

func PaddingRight(s string, l int) string {
	l -= len(s)
	if l <= 0 {
		return s
	}

	buf := new(bytes.Buffer)
	_, _ = buf.WriteString(s)

	sp := []byte(" ")
	for i := 0; i < l; i++ {
		buf.Write(sp)
	}

	return buf.String()
}
