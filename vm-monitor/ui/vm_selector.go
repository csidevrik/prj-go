// vm_selector.go - Pantalla de selección de VMs con checkbox en Bubbletea
package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/key"
)

type vmItem struct {
	name     string
	checked  bool
}

func (i vmItem) Title() string       { return i.name }
func (i vmItem) Description() string { return "" }
func (i vmItem) FilterValue() string { return i.name }

// Keys para atajos
var keys = struct {
	toggle key.Binding
	selectAll key.Binding
	execute key.Binding
	quit     key.Binding
}{
	toggle: key.NewBinding(key.WithKeys("space"), key.WithHelp("space", "marcar/desmarcar")),
	selectAll: key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "seleccionar todos")),
	execute: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "apagar seleccionadas")),
	quit: key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "salir")),
}

type selectorModel struct {
	list      list.Model
	loading   bool
	selection map[int]bool
	quit      bool
	ready     bool
	output    string
}

func NewSelectorModel(vms []string) selectorModel {
	items := make([]list.Item, len(vms))
	for i, name := range vms {
		items[i] = vmItem{name: name, checked: false}
	}
	l := list.New(items, list.NewDefaultDelegate(), 50, 15)
	l.Title = "Selecciona las VMs que deseas apagar"
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.toggle, keys.selectAll, keys.execute, keys.quit}
	}
	return selectorModel{
		list:      l,
		selection: make(map[int]bool),
	}
}

func (m selectorModel) Init() tea.Cmd {
	return nil
}

func (m selectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.quit):
			m.quit = true
			return m, tea.Quit
		case key.Matches(msg, keys.toggle):
			i := m.list.Index()
			m.selection[i] = !m.selection[i]
		case key.Matches(msg, keys.selectAll):
			for i := 0; i < len(m.list.Items()); i++ {
				m.selection[i] = true
			}
		case key.Matches(msg, keys.execute):
			var selected []string
			for i, sel := range m.selection {
				if sel {
					selected = append(selected, m.list.Items()[i].(vmItem).name)
				}
			}
			m.output = strings.Join(selected, ", ")
			m.ready = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-2)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m selectorModel) View() string {
	if m.ready {
		return fmt.Sprintf("Apagando VMs seleccionadas: %s\n", m.output)
	}
	return m.list.View()
}

func RunSelector(vmNames []string) ([]string, error) {
	m := NewSelectorModel(vmNames)
	finalModel, err := tea.NewProgram(m).Run()
	if err != nil {
		return nil, err
	}
	model := finalModel.(selectorModel)
	if !model.ready {
		return nil, nil
	}
	return strings.Split(model.output, ", "), nil
}
