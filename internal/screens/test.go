package screens

import (
	"strconv"
	"time"
	tea "charm.land/bubbletea/v2"

	"github.com/Adarshrai24/ttyper/internal/data"
	"github.com/Adarshrai24/ttyper/internal/test"
)

type TestModel struct {
	paragraph []rune
	typed     []rune

	cursor   int
	mistakes int

	started bool

	Duration int
	Remaining int
	
	Finished bool
}

func (m TestModel) TotalChars() int {
	return len(m.typed)
}

func (m TestModel) Mistakes() int {
	return m.mistakes
}

func (m TestModel) Result() test.Result {
	return test.Calculate(
		len(m.typed),
		m.mistakes,
		m.Duration,
	)
}

type TickMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}

func NewTest() TestModel {
	return TestModel{}
}

func (m *TestModel) Reset(duration int) {
	m.Duration = duration

	passages := data.RandomParagraphPick()
	m.paragraph = []rune(passages.Passages[passages.Current])

	m.typed = []rune{}
	m.cursor = 0
	m.mistakes = 0
	m.started = false
	m.Finished = false
	m.Remaining = duration
}

func (m TestModel) Update(msg tea.Msg) (TestModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case TickMsg:
		if m.started && !m.Finished {
			m.Remaining--

			if m.Remaining <= 0 {
				m.started = false
				m.Finished = true
				return m, nil
			}

			return m, tick()
		}

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			m.started = false
			m.Finished = true
			return m, nil

		case "backspace":
			if len(m.typed) > 0 {
				m.typed = m.typed[:len(m.typed)-1]

				if m.cursor > 0 {
					m.cursor--
				}
			}

		default:
			if len(msg.Text) == 0 {
				return m, nil
			}

			// Start timer on first keypress, but don't consume the key.
			if !m.started {
				m.started = true
				cmd = tick()
			}

			if m.cursor >= len(m.paragraph) {
				return m, cmd
			}

			r := []rune(msg.Text)[0]

			m.typed = append(m.typed, r)

			if r != m.paragraph[m.cursor] {
				m.mistakes++
			}

			m.cursor++

			if m.cursor == len(m.paragraph) {
				m.started = false
				m.Finished = true
			}
		}
	}

	return m, cmd
}

func (m TestModel) View() tea.View {
	s := "Typing Test\n\n"

	s += "Duration : "
	s += strconv.Itoa(m.Duration)
	s += " seconds\n\n"
	s += "Time Left : "
	s += strconv.Itoa(m.Remaining)
	s += " seconds\n\n"
	s += string(m.paragraph)

	s += "\n\nYour Input:\n"
	s += string(m.typed)

	s += "\n\nEsc - Finish"
	s += "\nCtrl+C - Quit"

	return tea.NewView(s)
}
