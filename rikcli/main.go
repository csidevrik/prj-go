package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/csidevrik/rikcli/internal/app"
)

func main() {
	if _, err := tea.NewProgram(app.NewModel()).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
