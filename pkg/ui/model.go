package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/clive/pkg/config"
	"github.com/koki-develop/clive/pkg/ttyd"
)

type Model struct {
	configFile string
	config     *config.Config
	ttyd       *ttyd.Ttyd

	err error

	spinner spinner.Model
}

var _ tea.Model = (*Model)(nil)

func New(configFile string) *Model {
	return &Model{
		configFile: configFile,
		spinner:    spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(styleSpinner)),
	}
}

func (m *Model) Err() error {
	return m.err
}

func (m *Model) Close() error {
	if m.ttyd == nil {
		return nil
	}

	if err := m.ttyd.Command.Process.Kill(); err != nil {
		return err
	}

	return nil
}
