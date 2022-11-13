package ui

import "fmt"

// TODO: implement
func (m *Model) View() string {
	if m.err != nil {
		return m.errView()
	}

	if m.config == nil {
		return m.loadingConfigView()
	}

	return m.loadingConfigView()
}

func (m *Model) errView() string {
	return "error" // TODO: implement
}

func (m *Model) loadingConfigView() string {
	return fmt.Sprintf("%s Loading config", m.spinner.View())
}
