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
	browser *rod.Browser
	page    *rod.Page
}

type actionDoneMsg struct{}

type pauseActionMsg struct{}

type pauseBeforeQuitMsg struct{}

type errMsg struct {
	err error
}
