package cmd

import (
	"fmt"
	"os/exec"
)

func ttyd(port int) *exec.Cmd {
	args := []string{
		fmt.Sprintf("--port=%d", port),
		"-t", "rendererType=canvas",
		"-t", "disableResizeOverlay=true",
		"-t", "cursorBlink=true",
		"-t", "customGlyphs=true",
		"--",
		"bash", "--login",
	}

	return exec.Command("ttyd", args...)
}
