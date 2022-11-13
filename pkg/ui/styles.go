package ui

import "github.com/charmbracelet/lipgloss"

var (
	// colors
	colorMain = "#ff00ff"

	// styles
	styleSpinner      = lipgloss.NewStyle().Foreground(lipgloss.Color(colorMain))
	styleActive       = lipgloss.NewStyle().Bold(true)
	styleActionHeader = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(lipgloss.Color(colorMain))
)
