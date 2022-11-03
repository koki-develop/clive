package cmd

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func launchBrowser() (*rod.Browser, error) {
	path, _ := launcher.LookPath()

	u, err := launcher.New().
		Leakless(true).
		Headless(false).
		Bin(path).
		Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().ControlURL(u)
	if err := browser.Connect(); err != nil {
		return nil, err
	}

	return browser, nil
}
