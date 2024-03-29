package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-rod/rod"
	"github.com/koki-develop/clive/internal/config"
	"github.com/koki-develop/clive/internal/net"
	"github.com/koki-develop/clive/internal/styles"
	"github.com/koki-develop/clive/internal/ttyd"
)

type Model struct {
	err error

	configFile string
	config     *config.Config

	ttyd *ttyd.Ttyd
	page *rod.Page

	currentActionIndex int
	pausing            bool
	quitting           bool

	spinner spinner.Model

	netListener net.IListener
}

var _ tea.Model = (*Model)(nil)

func New(configFile string) *Model {
	return &Model{
		configFile:         configFile,
		currentActionIndex: 0,
		spinner:            spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(styles.StyleSpinner)),
		netListener:        net.NewListener(),
	}
}

func (m *Model) Err() error {
	return m.err
}

func (m *Model) Close() error {
	if m.ttyd == nil {
		return nil
	}

	if err := m.ttyd.Close(); err != nil {
		return err
	}

	return nil
}

func (m *Model) running() bool {
	return m.page != nil
}
