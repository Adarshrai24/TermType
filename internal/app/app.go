package app

import (
	"database/sql"
	"log"

	tea "charm.land/bubbletea/v2"

	"github.com/Adarshrai24/ttyper/internal/screens"
	"github.com/Adarshrai24/ttyper/internal/storage"
)

type Screen int

const (
	MainScreen Screen = iota
	TimeScreen
	TestScreen
	ResultScreen
	HistoryScreen
	StatsScreen
)

type Model struct {
	db *sql.DB

	screen Screen

	menu    screens.MenuModel
	time    screens.TimeModel
	test    screens.TestModel
	result  screens.ResultModel
	history screens.HistoryModel
	stats   screens.StatsModel
}

func New() Model {
	db, err := storage.Open()
	if err != nil {
		log.Fatal(err)
	}

	return Model{
		db: db,

		screen: MainScreen,

		menu:    screens.NewMenu(),
		time:    screens.NewTime(),
		test:    screens.NewTest(),
		result:  screens.NewResult(),
		history: screens.NewHistory(),
		stats:   screens.NewStats(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.screen {

	case MainScreen:
		var cmd tea.Cmd

		m.menu, cmd = m.menu.Update(msg)

		if m.menu.StartPressed {
			m.menu.StartPressed = false
			m.screen = TimeScreen
		}

		if m.menu.HistoryPressed {
			m.menu.HistoryPressed = false

			_ = m.history.Load(m.db)

			m.screen = HistoryScreen
		}

		if m.menu.StatsPressed {
			m.menu.StatsPressed = false

			_ = m.stats.Load(m.db)

			m.screen = StatsScreen
		}

		return m, cmd

	case TimeScreen:
		var cmd tea.Cmd

		m.time, cmd = m.time.Update(msg)

		if m.time.Back {
			m.time.Back = false
			m.screen = MainScreen
		}

		if m.time.Start {
			m.time.Start = false
			m.test.Reset(m.time.Duration)
			m.screen = TestScreen
		}

		return m, cmd

	case TestScreen:
		var cmd tea.Cmd

		m.test, cmd = m.test.Update(msg)

		if m.test.Finished {
			m.test.Finished = false

			result := m.test.Result()
			err := storage.SaveSession(
				m.db,
				result,
				m.test.Duration,
			)
			if err != nil {
				return m, tea.Quit
			}

			m.result.SetResult(
				result,
				m.test.Duration,
			)

			m.screen = ResultScreen
		}

		return m, cmd

	case ResultScreen:
		var cmd tea.Cmd

		m.result, cmd = m.result.Update(msg)

		if m.result.Back {
			m.result.Back = false
			m.screen = MainScreen
		}

		return m, cmd

	case HistoryScreen:
		var cmd tea.Cmd

		m.history, cmd = m.history.Update(msg)

		if m.history.Back {
			m.history.Back = false
			m.screen = MainScreen
		}

		return m, cmd

	case StatsScreen:
		var cmd tea.Cmd

		m.stats, cmd = m.stats.Update(msg)

		if m.stats.Back {
			m.stats.Back = false
			m.screen = MainScreen
		}

		return m, cmd
	}

	return m, nil
}

func (m Model) View() tea.View {
	switch m.screen {

	case MainScreen:
		return m.menu.View()

	case TimeScreen:
		return m.time.View()

	case TestScreen:
		return m.test.View()

	case ResultScreen:
		return m.result.View()

	case HistoryScreen:
		return m.history.View()

	case StatsScreen:
		return m.stats.View()

	default:
		return tea.NewView("")
	}
}
