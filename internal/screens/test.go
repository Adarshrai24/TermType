package screens

import (
	"strconv"
	"time"
	"strings"
	tea "charm.land/bubbletea/v2"
	"github.com/muesli/reflow/wordwrap"
	"github.com/Adarshrai24/TermType/internal/data"
	"github.com/Adarshrai24/TermType/internal/test"
)

type TestModel struct {
	passages []string
	paragraph string
	wrappedLines []string 
	currentLine int 
	line []rune
	typed     []rune
	current int
	cursor   int
	mistakes int
	totalChars int

	Width int
	started bool

	Duration int
	Remaining int
	
	Finished bool
}

func (m TestModel) Mistakes() int {
	return m.mistakes
}

func (m TestModel) Result() test.Result {
	return test.Calculate(
		m.totalChars,
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
	m.passages = passages.Passages
	width := m.Width 
	if width == 0 {
		width = 80
	}
	m.paragraph = passages.Passages[passages.Current]
	m.wrappedLines = strings.Split(wordwrap.String(m.paragraph, width-4), "\n",)
	m.currentLine = 0
	m.line = []rune(m.wrappedLines[0])
	m.current = passages.Current
	m.totalChars = 0
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

			r := []rune(msg.Text)[0]

			m.typed = append(m.typed, r)
			m.totalChars++

			if r != m.line[m.cursor] {
				m.mistakes++
			}

			m.cursor++	
			if m.cursor == len(m.line) {
				m.currentLine++
				if m.currentLine >= len(m.wrappedLines) {
					m.current = (m.current + 1) % len(m.passages)
					width := m.Width
					if width == 0 {
						width = 80
					}
					m.paragraph = m.passages[m.current]
					m.wrappedLines = strings.Split(wordwrap.String(m.paragraph, width-4), "\n",)
					m.currentLine = 0
				}
				m.line = []rune(m.wrappedLines[m.currentLine])
				m.cursor = 0
				m.typed = nil
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
	}

	return m, cmd
}

func (m TestModel) View() tea.View {
	width := m.Width
	if width == 0 {
		width = 80
	}

	s := "Typing Test\n\n"

	s += "Duration : "
	s += strconv.Itoa(m.Duration)
	s += " seconds\n\n"
	s += "Time Left : "
	s += strconv.Itoa(m.Remaining)
	s += " seconds\n\n"
	s += string(m.line)

	s += "\n\n> "
	s += string(m.typed)

	s += "\n\nEsc - Finish"
	s += "\nCtrl+C - Quit"

	return tea.NewView(s)
}
