package main

import (
	"fmt"
	"os"
    "strings"

	"documentyzer/internal/scanner"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: documentyzer <path-del-repo>")
		os.Exit(1)
	}

	repoPath := os.Args[1]

	folders, err := scanner.ScanScriptsDir(repoPath)
	if err != nil {
		fmt.Printf("Error escaneando repo: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%-25s %-10s %-10s %-10s\n", "FOLDER", "META", "README", "INSTALL")
	fmt.Println(strings.Repeat("-", 60))
	for _, f := range folders {
		fmt.Printf("%-25s %-10v %-10v %-10v\n",
			f.Name,
			boolIcon(f.HasMeta),
			boolIcon(f.HasReadme),
			boolIcon(f.HasInstall),
		)
	}
}

func boolIcon(b bool) string {
	if b {
		return "✅"
	}
	return "❌"
}