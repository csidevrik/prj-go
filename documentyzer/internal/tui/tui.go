package tui

import (
	"fmt"
	"os"
	"strings"

	"documentyzer/internal/git"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	branches    []string
	err         error
	readmes     map[string]string
	showReadmes bool
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

type readmesMsg struct {
	readmes map[string]string
	err     error
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case branchMsg:
		m.branches = msg.branches
		m.err = msg.err
		return m, nil
	case readmesMsg:
		m.readmes = msg.readmes
		m.err = msg.err
		m.showReadmes = true
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			return m, func() tea.Msg {
				err := git.CheckoutAllRemoteBranches()
				branches, berr := git.ListBranches()
				if err == nil {
					err = berr
				}
				return branchMsg{branches, err}
			}
		case "d":
			return m, func() tea.Msg {
				err := git.DeleteAllLocalBranchesExceptMain()
				branches, berr := git.ListBranches()
				if err == nil {
					err = berr
				}
				return branchMsg{branches, err}
			}
		case "r":
			return m, func() tea.Msg {
				readmes, err := git.ReadReadmesFromBranches(m.branches)
				return readmesMsg{readmes, err}
			}
		case "b":
			m.showReadmes = false
			return m, nil
		case "q", "ctrl+c":
			return m, tea.Quit
		}
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\nPresiona q para salir.", m.err)
	}
	if m.showReadmes {
		var sb strings.Builder
		sb.WriteString("README.md de cada rama:\n\n")
		for branch, content := range m.readmes {
			sb.WriteString(fmt.Sprintf("== %s ==\n", branch))
			sb.WriteString(content)
			sb.WriteString("\n----------------------\n")
		}
		sb.WriteString("\nPresiona 'b' para volver o 'q' para salir.")
		return sb.String()
	}
	// Default view when not showing readmes
	return "Presiona 'l' para listar ramas, 'r' para leer README.md, 'd' para borrar ramas locales, 'q' para salir."
}

func Start() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
