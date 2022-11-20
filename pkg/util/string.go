package util

import (
	"bytes"
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

func String(v string) *string {
	return &v
}

func Contains(slice []string, r string) bool {
	for _, l := range slice {
		if l == r {
			return true
		}
	}
	return false
}

func Border(str string, style lipgloss.Style) string {
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
