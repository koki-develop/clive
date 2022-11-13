package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-rod/rod"
	"github.com/koki-develop/clive/pkg/config"
	"github.com/koki-develop/clive/pkg/ttyd"
)

type loadConfigMsg struct{ config *config.Config }
type startTtydMsg struct{ ttyd *ttyd.Ttyd }
type openMsg struct{ page *rod.Page }
type errMsg struct{ err error }

// TODO: implement
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// spinner
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	// key
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	// events
	case loadConfigMsg:
		m.config = msg.config
		return m, m.startTtyd
	case startTtydMsg:
		m.ttyd = msg.ttyd
		return m, m.open
	case openMsg:
		m.page = msg.page
	}

	return m, nil
}

func (m *Model) loadConfig() tea.Msg {
	cfg, err := config.Load(m.configFile)
	if err != nil {
		return errMsg{err}
	}

	return loadConfigMsg{cfg}
}

func (m *Model) startTtyd() tea.Msg {
	ttyd, err := ttyd.NewTtyd(m.config.Settings.LoginCommand)
	if err != nil {
		return errMsg{err}
	}

	if err := ttyd.Command.Start(); err != nil {
		return errMsg{err}
	}

	return startTtydMsg{ttyd}
}

func (m *Model) open() tea.Msg {
	page, err := openPage(m.config, m.ttyd.Port)
	if err != nil {
		return errMsg{err}
	}

	return openMsg{page}
}
