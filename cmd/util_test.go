package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
