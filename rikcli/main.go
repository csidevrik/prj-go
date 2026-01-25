package main

import (
	"fmt"
	"os"

	"rikcli/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(app.NewModel()).Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
