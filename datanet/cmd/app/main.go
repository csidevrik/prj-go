package main

import (
	"datanet/internal/macparser"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input  string
	output string
	err    error
}

// Init implements tea.Model's Init method.
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			res, err := macparser.DetectAndNormalize(m.input)
			if err != nil {
				m.err = err
				m.output = ""
			} else {
				m.err = nil
				m.output = fmt.Sprintf("✅ MAC: %s\nTipo: %s\nFabricante: %s", res.LinuxFormat, res.MACType, res.OUI)
			}
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		default:
			m.input += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString("🔍 Ingresa una dirección MAC:\n> " + m.input + "\n\n")
	if m.err != nil {
		b.WriteString("❌ Error: " + m.err.Error())
	} else if m.output != "" {
		b.WriteString(m.output)
	}
	return b.String()
}

// mac := " b022.7aea.bb6-d   "
// info, err := macparser.DetectAndNormalize(mac)
//
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
// fmt.Println("Tipo:", info.MACType)
// fmt.Println("Formato Linux: \n", info.LinuxFormat)
// fmt.Println("Formato Huawei: \n", info.HuaweiFormat)
// fmt.Println("Formato Cisco: \n", info.CiscoFormat)
func main() {
	p := tea.NewProgram(model{})
	_, err := p.Run()
	if err != nil {
		fmt.Println("❌ Error al ejecutar:", err)
		return
	}

	fmt.Println("\n👋 Gracias por usar datanet. ¡Hasta luego!\n")
}
