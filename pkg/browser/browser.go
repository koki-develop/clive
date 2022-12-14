package browser

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type BrowserConfig struct {
	Bin      *string
	URL      string
	Headless bool
}

func Open(cfg *BrowserConfig) (*rod.Page, error) {
	b, err := launchBrowser(cfg)
	if err != nil {
		return nil, err
	}

	p, err := b.Page(proto.TargetCreateTarget{URL: cfg.URL})
	if err != nil {
		return nil, err
	}
	if err := p.WaitIdle(time.Minute); err != nil {
		return nil, err
	}

	return p, nil
}

func launchBrowser(cfg *BrowserConfig) (*rod.Browser, error) {
	path, _ := launcher.LookPath()
	if cfg.Bin != nil {
		path = *cfg.Bin
	}

	u, err := launcher.New().
		Leakless(true).
		Headless(cfg.Headless).
		Bin(path).
		Launch()
	if err != nil {
		return nil, err
	}

	b := rod.New().ControlURL(u).NoDefaultDevice()
	if err := b.Connect(); err != nil {
		return nil, err
	}

	return b, nil
}
