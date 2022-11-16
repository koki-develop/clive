package ui

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/koki-develop/clive/pkg/browser"
	"github.com/koki-develop/clive/pkg/config"
)

func openPage(cfg *config.Config, port int) (*rod.Page, error) {
	url := fmt.Sprintf("http://localhost:%d", port)
	p, err := browser.Open(cfg.Settings.BrowserBin, url)
	if err != nil {
		return nil, err
	}

	if err := setupPage(cfg, p); err != nil {
		return nil, err
	}

	return p, nil
}

func setupPage(cfg *config.Config, page *rod.Page) error {
	// font family
	if cfg.Settings.FontFamily != nil {
		if _, err := page.Eval(fmt.Sprintf("() => term.options.fontFamily = '%s'", *cfg.Settings.FontFamily)); err != nil {
			return err
		}
	}

	// font size
	if _, err := page.Eval(fmt.Sprintf("() => term.options.fontSize = %d", cfg.Settings.FontSize)); err != nil {
		return err
	}
	if _, err := page.Eval("term.fit"); err != nil {
		return err
	}

	return nil
}
