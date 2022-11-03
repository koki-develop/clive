package cmd

import (
	"os/exec"

	"github.com/go-rod/rod"
)

type configLoadedMsg struct {
	config *config
}

type ttydStartedMsg struct {
	port int
	ttyd *exec.Cmd
}

type browserLaunchedMsg struct {
	browser *rod.Browser
	page    *rod.Page
}

type actionDoneMsg struct{}

type pauseActionMsg struct{}

type pauseBeforeQuitMsg struct{}

type errMsg struct {
	err error
}
