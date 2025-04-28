package main

import (
	"fmt"
	"os"

	"github.com/csidevrik/perfilizer/cmd"
)

func main() {
	// Ejecuta el comando raíz y maneja errores
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
