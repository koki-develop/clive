package ttyd

import (
	"fmt"
	"net"
	"os/exec"
)

type Ttyd struct {
	Args []string

	port    *int
	command *exec.Cmd
}

func New(args []string) *Ttyd {
	return &Ttyd{
		Args: args,
	}
}

func (ttyd *Ttyd) Start() error {
	port, err := randomUnusedPort()
	if err != nil {
		return err
	}

	args := []string{
		fmt.Sprintf("--port=%d", port),
		// See: https://github.com/tsl0922/ttyd/blob/main/man/ttyd.man.md#client-optoins
		"-t", "rendererType=canvas",
		"-t", "disableResizeOverlay=true",
		"-t", "cursorBlink=true",
		"--",
	}
	args = append(args, ttyd.Args...)

	ttyd.port = &port
	ttyd.command = exec.Command("ttyd", args...)
	if err := ttyd.command.Start(); err != nil {
		return err
	}

	return nil
}

func (ttyd *Ttyd) Port() *int {
	return ttyd.port
}

func (ttyd *Ttyd) Close() error {
	if ttyd.command == nil {
		return nil
	}

	if err := ttyd.command.Process.Kill(); err != nil {
		return err
	}

	return nil
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
