package styles

import "github.com/charmbracelet/lipgloss"

var (
	// colors
	ColorMain = lipgloss.Color("#ff00ff")
	ColorNote = lipgloss.Color("#0000ff")
	ColorErr  = lipgloss.Color("#ff0000")

	// styles
	StyleSpinner = lipgloss.NewStyle().Foreground(ColorMain)

	StyleActionHeader = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(ColorMain)
	StyleNoteHeader   = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(ColorNote)
	StyleErrorHeader  = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(ColorErr)

	StyleNotificationBorder = lipgloss.NewStyle().Foreground(ColorMain)
	StyleNotificationText   = lipgloss.NewStyle().Bold(true)

	StyleActive    = lipgloss.NewStyle().Bold(true)
	StyleDone      = lipgloss.NewStyle().Faint(true)
	StyleTruncated = lipgloss.NewStyle().Faint(true)
	StyleLink      = lipgloss.NewStyle().Underline(true)
)
