package cmd

import (
	"fmt"
	"os/exec"
)

type ttyd struct {
	Port    int
	Command *exec.Cmd
}

func newTtyd(cmd []string) (*ttyd, error) {
	port, err := randomUnusedPort()
	if err != nil {
		return nil, err
	}

	args := []string{
		fmt.Sprintf("--port=%d", port),
		"-t", "rendererType=canvas",
		"-t", "disableResizeOverlay=true",
		"-t", "cursorBlink=true",
		"-t", "customGlyphs=true",
		"--",
	}
	args = append(args, cmd...)

	return &ttyd{
		Port:    port,
		Command: exec.Command("ttyd", args...),
	}, nil
}
