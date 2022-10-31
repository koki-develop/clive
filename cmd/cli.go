package cmd

import (
	"fmt"
	"net"
	"os/exec"
)

func randomUnusedPort() (int, error) {
	addr, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	if err = addr.Close(); err != nil {
		return 0, err
	}

	return addr.Addr().(*net.TCPAddr).Port, nil
}

func ttyd(port int) *exec.Cmd {
	args := []string{
		fmt.Sprintf("--port=%d", port),
		"-t", "rendererType=canvas",
		"-t", "disableResizeOverlay=true",
		"-t", "cursorBlink=true",
		"-t", "customGlyphs=true",
		"bash", "--login",
	}

	return exec.Command("ttyd", args...)
}
