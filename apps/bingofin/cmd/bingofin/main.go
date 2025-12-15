package main

import (
	"fmt"
	"os"

	"github.com/csidevrik/bingofin/internal/app"
)

var (
	version = "1.0.0"
	commit  = "iniciando"
	date    = "2024-12-14"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v", "version":
			fmt.Printf("bingofin %s (commit=%s, date=%s)\n", version, commit, date)
			return
		case "--help", "-h", "help":
			fmt.Println("Usage:")
			fmt.Println("  bingofin --version")
			fmt.Println("  bingofin")
			return
		}
		fmt.Println("bingofin starting...")
	}
	app.Run()
}
