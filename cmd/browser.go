package cmd

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/koki-develop/clive/pkg/config"
)

func launchBrowser(cfg *config.Config) (*rod.Browser, error) {
	path, _ := launcher.LookPath()
	if cfg.Settings.BrowserBin != nil {
		path = *cfg.Settings.BrowserBin
	}

	u, err := launcher.New().
		Leakless(true).
		Headless(false).
		Bin(path).
		Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().ControlURL(u).NoDefaultDevice()
	if err := browser.Connect(); err != nil {
		return nil, err
	}

	return browser, nil
}
