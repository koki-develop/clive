package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/clive/pkg/config"
)

type Model struct {
	configFile string
	config     *config.Config

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

// TODO: implement
func (m *Model) Close() error {
	return nil
}
