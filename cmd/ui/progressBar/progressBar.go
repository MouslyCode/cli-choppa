package progressbar

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var body = lipgloss.NewStyle().Padding(1, 2)

type model struct {
	value int
	width int
	state tea.ProgressBarState
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View
