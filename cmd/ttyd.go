package cmd

import (
	"fmt"
	"os/exec"
)

func ttyd(port int, cmd []string) *exec.Cmd {
	args := []string{
		fmt.Sprintf("--port=%d", port),
		"-t", "rendererType=canvas",
		"-t", "disableResizeOverlay=true",
		"-t", "cursorBlink=true",
		"-t", "customGlyphs=true",
		"--",
	}
	args = append(args, cmd...)

	return exec.Command("ttyd", args...)
}
