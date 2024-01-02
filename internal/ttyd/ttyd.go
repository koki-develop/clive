package ttyd

import (
	"fmt"
	"os/exec"
)

type Ttyd struct {
	Port int

	command *exec.Cmd
}

func New(cmd []string, port int) *Ttyd {
	args := append([]string{
		fmt.Sprintf("--port=%d", port),
		// See: https://github.com/tsl0922/ttyd/blob/main/man/ttyd.man.md#client-optoins
		"-t", "titleFixed=cLive",
		"-t", "rendererType=canvas",
		"-t", "disableResizeOverlay=true",
		"-t", "cursorBlink=true",
		"--once",
		"--writable",
		"--",
	}, cmd...)

	return &Ttyd{
		Port:    port,
		command: exec.Command("ttyd", args...),
	}
}

func (t *Ttyd) Start() error {
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
