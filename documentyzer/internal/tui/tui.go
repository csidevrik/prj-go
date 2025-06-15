package tui

import (
	"fmt"
	"os"
	"strings"

	"../internal/git"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	branches []string
	err      error
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		branches, err := git.ListBranches()
		return branchMsg{branches, err}
	}
}

type branchMsg struct {
	branches []string
	err      error
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case branchMsg:
		m.branches = msg.branches
		m.err = msg.err
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPresiona q para salir.", m.err)
	}
	if len(m.branches) == 0 {
		return "Cargando ramas...\n"
	}
	return "Ramas del repositorio:\n" + strings.Join(m.branches, "\n") + "\n\nPresiona q para salir."
}

func Start() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
