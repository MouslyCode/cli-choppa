package timeChoice

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#890707")).
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA"))
	baseStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#331717")).
			Padding(0, 3).
			Margin(3)
	decideStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#331717")).
			Border(lipgloss.NormalBorder()).
			Padding(0, 3).
			Margin(3)
	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#890707")).
			Bold(true)

	bgContent = "Background Application Content\nLine 2 of the app layout\nLine 3 of the app layout"
	bgLayer   = lipgloss.NewLayer(bgContent).
			ID("background").
			X(0).
			Y(0).
			Z(0)
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel(choices []string, selected map[int]struct{}) model {
	return model{
		choices:  []string{"Pomodoro", "Short Break", "Long Break"},
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
	// The header
	s := "I repat myself when i under stress\n\n"

	// Iterate over our choices

	for i, choice := range m.choices {
		choice = baseStyle.Render(choice)
		// Is the cursor pointing at this choice?
		cursor := "" // no cursor
		if m.cursor == i {
			cursor = focusedStyle.Render(">") // cursor!
			if _, ok := m.selected[i]; !ok {
				choice = decideStyle.Render(choice)
			}
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
			choice = selectedStyle.Render(choice)
		}

		// Render the row
		s += fmt.Sprintf("%s%s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
