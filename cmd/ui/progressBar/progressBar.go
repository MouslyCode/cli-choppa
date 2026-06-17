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
