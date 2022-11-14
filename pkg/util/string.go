package util

import (
	"bytes"
	"strings"
	"unicode/utf8"
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
