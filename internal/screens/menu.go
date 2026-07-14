package screens 

import(
	tea "charm.land/bubbletea/v2"
)

type MenuModel struct {
	cursor int 
	choices []string 
	StartPressed bool 
	HistoryPressed bool
	StatsPressed bool
}

func NewMenu() MenuModel {
	return MenuModel {
		choices: []string{"Start Test", "History", "Statistics",},
	}
}

func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	switch msg := msg.(type) {
	
	case tea.KeyPressMsg:
		switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			
			case "down", "j": 
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			
			case "enter":
				switch m.cursor {
				case 0:
					m.StartPressed = true
				case 1:
					m.HistoryPressed = true
				case 2:
					m.StatsPressed = true
				}
			
			case "ctrl+c", "q":
				return m, tea.Quit
		}
	}
	return m, nil
}

func (m MenuModel) View() tea.View {
	s := "Welcome To TTYPER\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if i == m.cursor {
			cursor = ">"
		}

		s += cursor + " " + choice + "\n"
	}

	s += "\n↑/↓ to move • Enter to select • q to quit"

	return tea.NewView(s)
}
