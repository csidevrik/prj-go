package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"vm-monitor/internal"

	"github.com/charmbracelet/bubbles/textinput"
)

type screen int

const (
	login screen = iota
	menu
	result
)

type model struct {
	screen    screen
	vcenter   string
	username  string
	password  string
	choices   []string
	cursor    int
	output    string
	executed  bool
	inputs    []textinput.Model
	focus     int
	err       error
}

func initialModel() model {
	inputVC := textinput.New()
	inputVC.Placeholder = "vCenter IP o FQDN"
	inputVC.Focus()

	inputUser := textinput.New()
	inputUser.Placeholder = "Usuario (DOMINIO\\usuario)"

	inputPass := textinput.New()
	inputPass.Placeholder = "Contraseña"
	inputPass.EchoMode = textinput.EchoPassword
	inputPass.EchoCharacter = '•'

	return model{
		screen: login,
		choices: []string{
			"Consultar VMs encendidas",
			"Salir",
		},
		inputs: []textinput.Model{inputVC, inputUser, inputPass},
		focus:  0,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.screen {
	case login:
		return updateLogin(m, msg)
	case menu:
		return updateMenu(m, msg)
	default:
		return m, nil
	}
}

func updateLogin(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.focus < len(m.inputs)-1 {
				m.focus++
			} else {
				// Recoger valores
				m.vcenter = m.inputs[0].Value()
				m.username = m.inputs[1].Value()
				m.password = m.inputs[2].Value()
				m.screen = menu
				return m, nil
			}
		case "up":
			if m.focus > 0 {
				m.focus--
			}
		case "down":
			if m.focus < len(m.inputs)-1 {
				m.focus++
			}
		}
	}

	for i := range m.inputs {
		if i == m.focus {
			m.inputs[i].Focus()
		} else {
			m.inputs[i].Blur()
		}
		m.inputs[i], cmd = m.inputs[i].Update(msg)
	}
	return m, cmd
}

func updateMenu(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.choices[m.cursor] {
			case "Consultar VMs encendidas":
				command := `Get-VM | Where-Object {$_.PowerState -eq 'PoweredOn'} | Select Name, PowerState, VMHost`
				m.output = internal.ExecutePowerShellWithAuth(m.vcenter, m.username, m.password, command)
				m.executed = true
			case "Salir":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	switch m.screen {
	case login:
		b.WriteString("🔐 Conexión a vCenter\n\n")
		for i := range m.inputs {
			b.WriteString(m.inputs[i].View() + "\n")
		}
		b.WriteString("\nPresiona Enter para continuar\n")
	case menu:
		b.WriteString("🔍 Monitor de VMs encendidas en vCenter\n\n")
		for i, choice := range m.choices {
			cursor := "  "
			if i == m.cursor {
				cursor = "➤ "
			}
			b.WriteString(fmt.Sprintf("%s%s\n", cursor, choice))
		}
		if m.executed {
			b.WriteString("\n📋 Resultado:\n")
			b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Render(m.output))
			b.WriteString("\n\nPresiona q para salir.\n")
		}
	}
	return b.String()
}

func main() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
