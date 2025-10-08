package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	title  = lipgloss.NewStyle().Bold(true)
	hint   = lipgloss.NewStyle().Faint(true)
	tabOn  = lipgloss.NewStyle().Bold(true).Underline(true)
	tabOff = lipgloss.NewStyle().Faint(true)
)

func (m Model) View() string {
	header := title.Render("rikcli") + "  " +
		tabs(m.Route) + "\n\n"

	var body string
	switch m.Route {
	case RouteCounter:
		body = m.Counter.View()
	case RouteTimer:
		body = m.Timer.View()
	}

	footer := "\n" + hint.Render("[tab] cambiar pestaña  [q] salir")

	return header + body + footer + "\n"
}

func tabs(r Route) string {
	c := "Counter"
	t := "Timer"
	if r == RouteCounter {
		c = tabOn.Render(c)
		t = tabOff.Render(t)
	} else {
		c = tabOff.Render(c)
		t = tabOn.Render(t)
	}
	return fmt.Sprintf("%s  |  %s", c, t)
}
