package ui

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	err error
}

var _ tea.Model = (*Model)(nil)

func New() *Model {
	return &Model{}
}

func (m *Model) Err() error {
	return m.err
}

// TODO: implement
func (m *Model) Close() error {
	return nil
}
