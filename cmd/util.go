package cmd

import (
	"bytes"
)

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

func ptrString(v string) *string {
	return &v
}

func contains(slice []string, r string) bool {
	for _, l := range slice {
		if l == r {
			return true
		}
	}
	return false
}
