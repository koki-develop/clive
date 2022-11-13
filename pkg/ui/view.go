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

	if m.page == nil {
		return m.openingView()
	}

	return m.loadingConfigView()
}

func (m *Model) errView() string {
	return "error" // TODO: implement
}

func (m *Model) loadingConfigView() string {
	return fmt.Sprintf("%s Loading config", m.spinner.View())
}

func (m *Model) openingView() string {
	return fmt.Sprintf("%s Opening", m.spinner.View())
}
