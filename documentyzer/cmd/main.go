package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"documentyzer/internal/git"
	"documentyzer/internal/project"
	"documentyzer/internal/scanner"
	"documentyzer/internal/tui"
)

func main() {
	if len(os.Args) < 2 {
		printMainHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "help", "-h", "--help":
		printMainHelp()
	case "scanner":
		scannerCommand()
	case "tui":
		tuiCommand()
	case "branches":
		branchesCommand()
	case "init":
		initCommand()
	default:
		fmt.Printf("Comando desconocido: %s\n\n", command)
		printMainHelp()
		os.Exit(1)
	}
}

// COMANDO: scanner
func scannerCommand() {
	fs := flag.NewFlagSet("scanner", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println("Uso: documentyzer scanner <path-del-repo>")
		fmt.Println()
		fmt.Println("Escanea un repositorio y analiza la estructura de carpetas scripts.d")
		fmt.Println()
		fmt.Println("Argumentos:")
		fmt.Println("  <path-del-repo>   Ruta al repositorio que contiene la carpeta scripts.d")
	}

	fs.Parse(os.Args[2:])
	args := fs.Args()

	if len(args) == 0 {
		fs.Usage()
		os.Exit(1)
	}

	repoPath := args[0]

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

// COMANDO: tui
func tuiCommand() {
	fs := flag.NewFlagSet("tui", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println("Uso: documentyzer tui")
		fmt.Println()
		fmt.Println("Abre la interfaz interactiva (TUI) para gestionar ramas y documentación")
		fmt.Println()
		fmt.Println("Controles:")
		fmt.Println("  l   - Listar y hacer checkout de ramas remotas")
		fmt.Println("  d   - Borrar todas las ramas locales excepto main")
		fmt.Println("  r   - Leer README.md de las diferentes ramas")
		fmt.Println("  b   - Volver a la vista principal")
		fmt.Println("  q   - Salir")
	}

	fs.Parse(os.Args[2:])

	fmt.Println("Iniciando interfaz interactiva...")
	tui.Start()
}

// COMANDO: branches
func branchesCommand() {
	fs := flag.NewFlagSet("branches", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println("Uso: documentyzer branches <subcomando>")
		fmt.Println()
		fmt.Println("Gestiona ramas del repositorio")
		fmt.Println()
		fmt.Println("Subcomandos:")
		fmt.Println("  list     - Listar todas las ramas")
		fmt.Println("  checkout - Hacer checkout de todas las ramas remotas")
		fmt.Println("  clean    - Borrar todas las ramas locales excepto main")
		fmt.Println("  readmes  - Leer README.md de todas las ramas")
	}

	fs.Parse(os.Args[2:])
	args := fs.Args()

	if len(args) == 0 {
		fs.Usage()
		os.Exit(1)
	}

	subcommand := args[0]

	switch subcommand {
	case "list":
		branches, err := git.ListBranches()
		if err != nil {
			fmt.Printf("Error listando ramas: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Ramas disponibles:")
		for _, b := range branches {
			if b != "" {
				fmt.Printf("  - %s\n", b)
			}
		}

	case "checkout":
		fmt.Println("Haciendo checkout de todas las ramas remotas...")
		if err := git.CheckoutAllRemoteBranches(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Checkout completado")

	case "clean":
		fmt.Println("Borrando todas las ramas locales excepto main...")
		if err := git.DeleteAllLocalBranchesExceptMain(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Limpieza completada")

	case "readmes":
		branches, err := git.ListBranches()
		if err != nil {
			fmt.Printf("Error listando ramas: %v\n", err)
			os.Exit(1)
		}
		readmes, err := git.ReadReadmesFromBranches(branches)
		if err != nil {
			fmt.Printf("Error leyendo READMEs: %v\n", err)
			os.Exit(1)
		}
		for branch, content := range readmes {
			fmt.Printf("=== %s ===\n%s\n\n", branch, content)
		}

	default:
		fmt.Printf("Subcomando desconocido: %s\n\n", subcommand)
		fs.Usage()
		os.Exit(1)
	}
}

// COMANDO: init
func initCommand() {
	fs := flag.NewFlagSet("init", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println("Uso: documentyzer init <path>")
		fmt.Println()
		fmt.Println("Inicializa la estructura de documentación en una carpeta individual")
		fmt.Println()
		fmt.Println("Argumentos:")
		fmt.Println("  <path>   Ruta de la carpeta a inicializar (carpeta dentro de scripts.d)")
		fmt.Println()
		fmt.Println("Descripción:")
		fmt.Println("  Este comando crea la estructura de documentación en una carpeta individual,")
		fmt.Println("  NO en la raíz del repositorio. Se usa para cada proyecto/script dentro de scripts.d")
		fmt.Println()
		fmt.Println("Estructura que crea:")
		fmt.Println("  - meta.json      Metadatos del proyecto (nombre, descripción, versión, etc)")
		fmt.Println("  - README.md      Documentación del proyecto")
		fmt.Println("  - install.sh     Script de instalación")
		fmt.Println()
		fmt.Println("Ejemplos:")
		fmt.Println("  documentyzer init /path/to/repo/scripts.d/mi-proyecto")
		fmt.Println("  documentyzer init /home/user/prj-bash/scripts.d/install-utils")
		fmt.Println("  documentyzer init ./scripts.d/my-script")
	}

	fs.Parse(os.Args[2:])
	args := fs.Args()

	if len(args) == 0 {
		fs.Usage()
		os.Exit(1)
	}

	folderPath := args[0]

	// Inicializar estructura
	if err := project.InitProject(folderPath); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n✅ Estructura inicializada en: %s\n", folderPath)
}

// AUXILIARES
func boolIcon(b bool) string {
	if b {
		return "✅"
	}
	return "❌"
}

func printMainHelp() {
	fmt.Println("documentyzer v1.0 - Herramienta de análisis y documentación de repositorios")
	fmt.Println()
	fmt.Println("Uso:")
	fmt.Println("  documentyzer <comando> [opciones]")
	fmt.Println()
	fmt.Println("Comandos:")
	fmt.Println("  scanner <path>    Escanea un repositorio y analiza scripts.d")
	fmt.Println("  tui               Inicia interfaz interactiva para gestionar ramas")
	fmt.Println("  branches <sub>    Gestiona ramas (list, checkout, clean, readmes)")
	fmt.Println("  init <path>       Inicializa estructura en carpeta dentro de scripts.d")
	fmt.Println("  help              Muestra esta ayuda")
	fmt.Println()
	fmt.Println("Ejemplos:")
	fmt.Println("  # Escanear un repositorio completo")
	fmt.Println("  documentyzer scanner /path/to/repo")
	fmt.Println()
	fmt.Println("  # Gestionar ramas (interfaz interactiva)")
	fmt.Println("  documentyzer tui")
	fmt.Println()
	fmt.Println("  # Operaciones con ramas (CLI)")
	fmt.Println("  documentyzer branches list")
	fmt.Println("  documentyzer branches checkout")
	fmt.Println("  documentyzer branches clean")
	fmt.Println()
	fmt.Println("  # Inicializar carpeta individual dentro de scripts.d")
	fmt.Println("  documentyzer init /path/to/repo/scripts.d/mi-proyecto")
	fmt.Println()
	fmt.Println("  # Ver ayuda")
	fmt.Println("  documentyzer help")
}