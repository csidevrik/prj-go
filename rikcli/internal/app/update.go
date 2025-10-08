package app

import (
	tea "github.com/charmbracelet/bubbletea"
)

type navMsg Route // para cambiar de ruta

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			// alterna entre rutas
			if m.Route == RouteCounter {
				m.Route = RouteTimer
			} else {
				m.Route = RouteCounter
			}
			return m, nil
		}

	case navMsg:
		m.Route = Route(msg)
		return m, nil
	}

	// delega a la página activa
	switch m.Route {
	case RouteCounter:
		newCounter, cmd := m.Counter.Update(msg)
		m.Counter = newCounter
		return m, cmd
	case RouteTimer:
		newTimer, cmd := m.Timer.Update(msg)
		m.Timer = newTimer
		return m, cmd
	default:
		return m, nil
	}
}
