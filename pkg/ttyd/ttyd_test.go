package ttyd

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		cmd  []string
		port int
	}
	tests := []struct {
		args args
		want *Ttyd
	}{
		{
			args{[]string{"hoge", "fuga"}, 999},
			&Ttyd{
				command: exec.Command(
					"ttyd",
					"--port=999",
					"-t", "rendererType=canvas",
					"-t", "disableResizeOverlay=true",
					"-t", "cursorBlink=true",
					"--",
					"hoge", "fuga",
				),
				Port: 999,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := New(tt.args.cmd, tt.args.port)
			assert.Equal(t, tt.want, got)
		})
	}
}
