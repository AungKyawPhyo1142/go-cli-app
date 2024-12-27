package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices:  []string{"Buy carrots", "Wash dishes", "Buy milk"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// this is just saying when something happend, update the "ui"
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	// is it a key-pressed?
	case tea.KeyMsg:
		// ok, which key is it?
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
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

func (m model) View() string {
	// the header
	s := "What do we have to do today??\n\n"

	//iterate the choices we have
	for i, choice := range m.choices {
		// is the cursor pointing at this choice?
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected
		}

		// render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// this is footer
	s += "\nPres q to qui.\n"

	return s
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Finally, there is an error: %v", err)
		os.Exit(1)
	}

}
