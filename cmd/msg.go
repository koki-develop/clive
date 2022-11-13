package cmd

import (
	"github.com/go-rod/rod"
	"github.com/koki-develop/clive/pkg/config"
	"github.com/koki-develop/clive/pkg/ttyd"
)

type configLoadedMsg struct {
	config *config.Config
}

type ttydStartedMsg struct {
	Ttyd *ttyd.Ttyd
}

type browserLaunchedMsg struct {
	Browser *rod.Browser
}

type pageOpenedMsg struct {
	Page *rod.Page
}

type actionDoneMsg struct{}

type pauseActionMsg struct{}

type pauseBeforeQuitMsg struct{}

type errMsg struct {
	Err error
}
