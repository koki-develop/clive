package browser

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func Open(bin *string, url string) (*rod.Page, error) {
	b, err := launchBrowser(bin)
	if err != nil {
		return nil, err
	}

	p, err := b.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		return nil, err
	}
	if err := p.WaitIdle(time.Minute); err != nil {
		return nil, err
	}

	return p, nil
}

func launchBrowser(bin *string) (*rod.Browser, error) {
	path, _ := launcher.LookPath()
	if bin != nil {
		path = *bin
	}

	u, err := launcher.New().
		Leakless(true).
		Headless(false).
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
