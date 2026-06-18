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

	baseChoice = lipgloss.NewStyle().
			Padding(0, 1).
			MarginLeft(2).
			Background(dark)

	activeChoice = baseChoice.Background(primary)
)

type Selection struct {
	Choice string
}

func (s *Selection) Update(value string) {
	s.Choice = value
}

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}

	choice *Selection
}

func InitialModel(choices []string, selected map[int]struct{}, result *Selection) model {
	if choices == nil {
		choices = []string{" Pomodoro ", " Short Break ", " Long Break "}
	}
	if selected == nil {
		selected = make(map[int]struct{})
	}
	return model{
		choices:  choices,
		selected: selected,
		choice:   result,
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

		case "q", "ctrl+c":
			return m, tea.Quit

		case "left", "j":
			if m.cursor > 0 {
				m.cursor--
			}

		case "right", "k":
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

	// Content
	s := "\n"
	s += row + "\n\n"
	s += "\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
