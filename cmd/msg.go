package cmd

import (
	"github.com/go-rod/rod"
)

type configLoadedMsg struct {
	config *config
}

type ttydStartedMsg struct {
	Ttyd *ttyd
}

type browserLaunchedMsg struct {
	Browser *rod.Browser
	Page    *rod.Page
}

type actionDoneMsg struct{}

type pauseActionMsg struct{}

type pauseBeforeQuitMsg struct{}

type errMsg struct {
	Err error
}
