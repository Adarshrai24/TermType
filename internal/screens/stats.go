package screens

import (
	"database/sql"
	"fmt"

	tea "charm.land/bubbletea/v2"

	"github.com/Adarshrai24/ttyper/internal/storage"
)

type StatsModel struct {
	Average []storage.Stats
	Maximum []storage.Stats

	Back bool
}

func NewStats() StatsModel {
	return StatsModel{}
}

func (m *StatsModel) Load(db *sql.DB) error {
	durations := []int{15, 30, 60, 120}

	m.Average = nil
	m.Maximum = nil

	for _, d := range durations {
		avg, err := storage.AverageStats(db, d)
		if err != nil {
			return err
		}

		max, err := storage.MaximumStats(db, d)
		if err != nil {
			return err
		}

		m.Average = append(m.Average, avg)
		m.Maximum = append(m.Maximum, max)
	}

	return nil
}

func (m StatsModel) Update(msg tea.Msg) (StatsModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "enter", "esc":
			m.Back = true

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m StatsModel) View() tea.View {
	s := "Statistics\n\n"

	s += "Average Statistics\n\n"
	s += fmt.Sprintf("%-10s %-10s %-10s %-10s\n",
		"Duration", "GWPM", "NWPM", "Accuracy")

	s += "-----------------------------------------------\n"

	for _, stat := range m.Average {
		s += fmt.Sprintf("%-10d %-10.2f %-10.2f %-10.2f%%\n",
			stat.Duration,
			stat.GWPM,
			stat.NWPM,
			stat.Accuracy,
		)
	}

	s += "\n"

	s += "Maximum Statistics\n\n"
	s += fmt.Sprintf("%-10s %-10s %-10s %-10s\n",
		"Duration", "GWPM", "NWPM", "Accuracy")

	s += "-----------------------------------------------\n"

	for _, stat := range m.Maximum {
		s += fmt.Sprintf("%-10d %-10.2f %-10.2f %-10.2f%%\n",
			stat.Duration,
			stat.GWPM,
			stat.NWPM,
			stat.Accuracy,
		)
	}

	s += "\nPress Enter or Esc to return."

	return tea.NewView(s)
}
