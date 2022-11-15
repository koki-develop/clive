package ttyd

import (
	"fmt"
	"net"
	"os/exec"
)

type Ttyd struct {
	Args []string
	Port *int

	command *exec.Cmd
}

func New(args []string) *Ttyd {
	return &Ttyd{
		Args: args,
	}
}

func (t *Ttyd) Start() error {
	port, err := t.randomUnusedPort()
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
	args = append(args, t.Args...)

	t.Port = &port
	t.command = exec.Command("ttyd", args...)
	if err := t.command.Start(); err != nil {
		return err
	}

	return nil
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

func (t *Ttyd) randomUnusedPort() (int, error) {
	addr, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}

	if err = addr.Close(); err != nil {
		return 0, err
	}

	return addr.Addr().(*net.TCPAddr).Port, nil
}
