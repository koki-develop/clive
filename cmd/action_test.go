package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTypeAction(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		args    args
		want    *typeAction
		wantErr bool
	}{
		{
			args{
				map[string]interface{}{
					"type": "Hello World",
				},
			},
			&typeAction{
				Type:  "Hello World",
				Count: 1,
				Speed: 10,
			},
			false,
		},
		{
			args{
				map[string]interface{}{
					"type":  "Hello World",
					"count": 10,
					"speed": 500,
				},
			},
			&typeAction{
				Type:  "Hello World",
				Count: 10,
				Speed: 500,
			},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got, err := parseTypeAction(tt.args.m)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
