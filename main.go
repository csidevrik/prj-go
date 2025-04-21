package main

import (
	"fmt"
	"os"

	"<TU_MÓDULO>/perfilizer/cmd"
)

func main() {
	// Ejecuta el comando raíz y maneja errores
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
