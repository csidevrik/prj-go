package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/csidevrik/rikcli/internal/pages"
)

type Route int

const (
	RouteCounter Route = iota
	RouteTimer
)

type Model struct {
	Route   Route
	Counter pages.CounterModel
	Timer   pages.TimerModel
}

func NewModel() Model {
	return Model{
		Route:   RouteCounter,
		Counter: pages.NewCounter(),
		Timer:   pages.NewTimer(),
	}
}

func (m Model) Init() tea.Cmd {
	// Si alguna página necesita init, encadénalo aquí
	return m.Timer.Init() // Timer arranca con un tick
}
