package main

import tea "github.com/charmbracelet/bubbletea"

const (
	ListView uint = iota
	TitleView
	BodyView
)

type model struct {
	state uint
}

func NewModel() model {
	return model{
		state: ListView,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
