package ui

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.loadConfig,
	)
}
