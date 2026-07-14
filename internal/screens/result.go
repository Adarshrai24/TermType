package screens

import (
	"strconv"
	"github.com/Adarshrai24/ttyper/internal/test"
	tea "charm.land/bubbletea/v2"
)

type ResultModel struct {
	Result test.Result

	Duration int

	Back bool
}

func NewResult() ResultModel {
	return ResultModel{}
}

func (m *ResultModel) SetResult(result test.Result, duration int) {
	m.Result = result
	m.Duration = duration
}

func (m ResultModel) Update(msg tea.Msg) (ResultModel, tea.Cmd) {
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

func (m ResultModel) View() tea.View {
	s := "Test Results\n\n"

	s += "Duration : "
	s += strconv.Itoa(m.Duration)
	s += " seconds\n\n"

	s += "GWPM     : "
	s += strconv.FormatFloat(m.Result.GWPM, 'f', 2, 64)
	s += "\n"

	s += "NWPM     : "
	s += strconv.FormatFloat(m.Result.NWPM, 'f', 2, 64)
	s += "\n"

	s += "Accuracy : "
	s += strconv.FormatFloat(m.Result.Accuracy, 'f', 2, 64)
	s += "%\n\n"

	s += "Press Enter or Esc to return."

	return tea.NewView(s)
}
