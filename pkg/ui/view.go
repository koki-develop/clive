package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/koki-develop/clive/pkg/util"
)

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

	s := m.actionsView()
	if m.quitting {
		s += "\n\n" + m.quittingView()
	}

	return s
}

func (m *Model) errView() string {
	return styleErrorHeader.Render("Error") + "\n" + m.err.Error() + "\n\n" + m.quittingView()
}

func (m *Model) loadingConfigView() string {
	return fmt.Sprintf("%s Loading config", m.spinner.View())
}

func (m *Model) openingView() string {
	return fmt.Sprintf("%s Opening", m.spinner.View())
}

func (m *Model) actionsView() string {
	from := util.Max(0, m.currentActionIndex-3)
	show := 20
	digits := len(strconv.Itoa(len(m.config.Actions)))

	rows := []string{}
	for i, action := range m.config.Actions {
		if i < from && len(m.config.Actions)-i > show {
			continue
		}
		if i-from >= show {
			rows = append(rows, fmt.Sprintf("... %d more actions", len(m.config.Actions)-i))
			break
		}

		style := lipgloss.NewStyle()

		cursor := "  "
		if m.currentActionIndex > i {
			style = style.Faint(true)
		} else if m.currentActionIndex == i {
			style = style.Bold(true)
			if !m.quitting {
				if m.pausing {
					cursor = "> "
				} else {
					cursor = m.spinner.View()
				}
			}
		}

		num := util.PaddingRight(fmt.Sprintf("#%d", i+1), digits+1)
		rows = append(rows, fmt.Sprintf("%s %s%s", style.Render(num), cursor, style.Render(action.String())))
	}

	return styleActionHeader.Render("Actions") + "\n" + strings.Join(rows, "\n")
}

func (m *Model) quittingView() string {
	return styleActive.Render("Press enter to quit")
}
