package ttyd

import (
	"fmt"
	"net"
	"os/exec"
)

type Ttyd struct {
	Port    int
	Command *exec.Cmd
}

func NewTtyd(cmd []string) (*Ttyd, error) {
	port, err := randomUnusedPort()
	if err != nil {
		return nil, err
	}

	args := []string{
		fmt.Sprintf("--port=%d", port),
		// See: https://github.com/tsl0922/ttyd/blob/main/man/ttyd.man.md#client-optoins
		"-t", "rendererType=canvas",
		"-t", "disableResizeOverlay=true",
		"-t", "cursorBlink=true",
		"--",
	}
	args = append(args, cmd...)

	return &Ttyd{
		Port:    port,
		Command: exec.Command("ttyd", args...),
	}, nil
}

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
