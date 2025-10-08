package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type CounterModel struct {
	Count int
}

func NewCounter() CounterModel { return CounterModel{Count: 0} }

func (m CounterModel) Init() tea.Cmd { return nil }

func (m CounterModel) Update(msg tea.Msg) (CounterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "+", "k", "up":
			m.Count++
		case "-", "j", "down":
			m.Count--
		}
	}
	return m, nil
}

func (m CounterModel) View() string {
	return fmt.Sprintf("🧮 Contador: %d\n\n[+ ↑ k] suma   [- ↓ j] resta", m.Count)
}
