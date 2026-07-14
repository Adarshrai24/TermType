package screens

import (
	"database/sql"
	"fmt"
	tea "charm.land/bubbletea/v2"

	"github.com/Adarshrai24/TermType/internal/storage"
)

type HistoryModel struct {
	history []storage.Session

	Back bool
}

func NewHistory() HistoryModel {
	return HistoryModel{}
}

func (m *HistoryModel) Load(db *sql.DB) error {
	history, err := storage.GetHistory(db)
	if err != nil {
		return err
	}

	m.history = history
	return nil
}

func (m HistoryModel) Update(msg tea.Msg) (HistoryModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "esc", "enter":
			m.Back = true

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m HistoryModel) View() tea.View {
	s := "History\n\n"

	if len(m.history) == 0 {
		s += "No typing tests found.\n"
	} else {
		s += fmt.Sprintf("%-5s %-10s %-10s %-12s %-10s\n",
			"#", "GWPM", "NWPM", "Accuracy", "Duration")
		s += "---------------------------------------------------------\n"

		for i, h := range m.history {
			s += fmt.Sprintf("%-5d %-10.2f %-10.2f %-12.2f %-10d\n",
				i+1,
				h.GWPM,
				h.NWPM,
				h.Accuracy,
				h.Duration,
			)
		}
	}

	s += "\nPress Enter or Esc to return."

	return tea.NewView(s)
}
