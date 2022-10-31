package cmd

import (
	"errors"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func launchBrowser() (*rod.Browser, error) {
	path, ok := launcher.LookPath()
	if !ok {
		return nil, errors.New("no executable browser was found")
	}

	u, err := launcher.New().
		Leakless(true).
		Headless(false).
		Bin(path).
		Launch()
	if err != nil {
		return nil, err
	}

	return rod.New().ControlURL(u).MustConnect(), nil
}
