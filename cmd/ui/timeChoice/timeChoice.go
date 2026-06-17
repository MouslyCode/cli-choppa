package timeChoice

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	width       = 96
	ColumnWidth = 30
)

var (
	primary = lipgloss.Color("#890707")
	dark    = lipgloss.Color("#1f0101")
	grey    = lipgloss.Color("#151515")

	baseChoice = lipgloss.NewStyle().
			Padding(0, 1).
			MarginLeft(2).
			Background(dark)

	activeChoice = baseChoice.Background(primary)

	subTitle = lipgloss.NewStyle().
			Foreground(primary).
			Bold(true)
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel(choices []string, selected map[int]struct{}) model {
	return model{
		choices:  []string{" Pomodoro ", " Short Break ", " Long Break "},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the space bar toggle the selected state
		// for the item that the cursor is pointing at.
		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {

	rendered := make([]string, len(m.choices))
	for i, c := range m.choices {
		if i == m.cursor {
			rendered[i] = activeChoice.Render(c)
		} else {
			rendered[i] = baseChoice.Render(c)
		}
	}

	row := lipgloss.JoinHorizontal(lipgloss.Center, rendered...)
	row = lipgloss.PlaceHorizontal(width, lipgloss.Center, row)
	header := lipgloss.PlaceHorizontal(width, lipgloss.Center, subTitle.Render("I repat myself when i under stress"))
	footer := lipgloss.PlaceHorizontal(width, lipgloss.Center, "Press q to quit")

	// The header
	s := header + "\n\n"

	// Content
	s += row + "\n\n"

	// The footer
	s += "\n" + footer + "\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
