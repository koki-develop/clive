package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	err error

	spinner spinner.Model
}

var _ tea.Model = (*Model)(nil)

func New() *Model {
	return &Model{
		spinner: spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(styleSpinner)),
	}
}

func (m *Model) Err() error {
	return m.err
}

// TODO: implement
func (m *Model) Close() error {
	return nil
}
