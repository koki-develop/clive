package util

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/text"
)

func TruncateString(s string, l int) (string, bool) {
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

	return s, trunc
}

func PaddingRight(s string, l int) string {
	return text.Pad(s, l, ' ')
}

func String(v string) *string {
	return &v
}

func Border(str string, style lipgloss.Style) string {
	// NOTE: Don't use `lipgloss.Border“.
	//	See https://github.com/charmbracelet/lipgloss/issues/40 .
	lines := strings.Split(str, "\n")
	width := text.LongestLineLen(str)

	b := strings.Repeat("─", width+2)
	bt := style.Render(fmt.Sprintf("┌%s┐", b))
	bb := style.Render(fmt.Sprintf("└%s┘", b))

	rslt := []string{bt}
	for _, line := range lines {
		b := style.Render("│")
		rslt = append(rslt, fmt.Sprintf("%s %s %s", b, PaddingRight(line, width), b))
	}
	rslt = append(rslt, bb)

	return strings.Join(rslt, "\n")
}
