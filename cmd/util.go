package cmd

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

func truncateString(s string, l int) string {
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
		s += "..."
	}

	return s
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func paddingRight(s string, l int) string {
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

func ptr[T any](v T) *T {
	return &v
}

func contains[T comparable](slice []T, r T) bool {
	for _, l := range slice {
		if l == r {
			return true
		}
	}
	return false
}
