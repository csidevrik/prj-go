package pages

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TimerModel struct {
	Seconds int
	Running bool
}

type tickMsg time.Time

func NewTimer() TimerModel {
	return TimerModel{Running: true}
}

func (m TimerModel) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return tickMsg(t) })
}

func (m TimerModel) Update(msg tea.Msg) (TimerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			m.Running = !m.Running
			if m.Running {
				return m, tick()
			}
		case "r":
			m.Seconds = 0
		}
	case tickMsg:
		if m.Running {
			m.Seconds++
			return m, tick()
		}
	}
	return m, nil
}

func (m TimerModel) View() string {
	state := "⏸️ pausa"
	if m.Running {
		state = "▶️ play"
	}
	return fmt.Sprintf("⏱️  Tiempo: %ds (%s)\n\n[espacio] play/pause   [r] reset", m.Seconds, state)
}
