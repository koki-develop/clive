package styles

import "github.com/charmbracelet/lipgloss"

var (
	// colors
	ColorMain = lipgloss.Color("#ff00ff")
	ColorErr  = lipgloss.Color("#ff0000")

	// styles
	StyleSpinner      = lipgloss.NewStyle().Foreground(ColorMain)
	StyleActive       = lipgloss.NewStyle().Bold(true)
	StyleActionHeader = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(ColorMain)
	StyleErrorHeader  = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(ColorErr)
	StyleDone         = lipgloss.NewStyle().Faint(true)
	StyleTruncated    = lipgloss.NewStyle().Faint(true)
)
