package ui

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/koki-develop/clive/pkg/config"
)

func openPage(cfg *config.Config, port int) (*rod.Page, error) {
	browser, err := launchBrowser(cfg)
	if err != nil {
		return nil, err
	}

	page, err := browser.Page(proto.TargetCreateTarget{URL: fmt.Sprintf("http://localhost:%d", port)})
	if err != nil {
		return nil, err
	}
	if err := page.WaitIdle(time.Minute); err != nil {
		return nil, err
	}
	if err := setupPage(cfg, page); err != nil {
		return nil, err
	}

	return page, nil
}

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

func setupPage(cfg *config.Config, page *rod.Page) error {
	if cfg.Settings.FontFamily != nil {
		if _, err := page.Eval(fmt.Sprintf("() => term.options.fontFamily = '%s'", *cfg.Settings.FontFamily)); err != nil {
			return err
		}
	}
	if _, err := page.Eval(fmt.Sprintf("() => term.options.fontSize = %d", cfg.Settings.FontSize)); err != nil {
		return err
	}
	if _, err := page.Eval("term.fit"); err != nil {
		return err
	}

	return nil
}
