package screens

import (
	"strconv"
	tea "charm.land/bubbletea/v2"
)

type TimeModel struct {
	cursor int

	durations []int

	Duration int

	Start bool
	Back  bool
}

func NewTime() TimeModel {
	return TimeModel{
		durations: []int{
			15,
			30,
			60,
			120,
		},
	}
}

func (m TimeModel) Update(msg tea.Msg) (TimeModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.durations)-1 {
				m.cursor++
			}

		case "enter":
			m.Duration = m.durations[m.cursor]
			m.Start = true

		case "esc":
			m.Back = true

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m TimeModel) View() tea.View {
	s := "Select Duration\n\n"

	for i, t := range m.durations {
		if i == m.cursor {
			s += "> "
		} else {
			s += "  "
		}

		s += strconv.Itoa(t) + " seconds\n"
	}

	s += "\nEnter - Start Test"
	s += "\nEsc   - Back"

	return tea.NewView(s)
}

