package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	profileDirectory = "Profile 2"
	unsionURL        = "https://www.youtube.com/watch?v=LJaSBWHvOd0&list=PLwtmRZw8oM40isWfAIltx-0aH0_phi4QO"
	startDelay       = 5
)

func main() {
	// Configurar logging
	logFile, err := os.OpenFile("/tmp/tv-lnx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	log.Println("=== 📺 Unsion Launcher INICIADO ===")
	fmt.Println("=== 📺 Unsion Launcher para Mamá Elvia ===")

	// Log de variables de entorno
	log.Printf("DISPLAY: %s\n", os.Getenv("DISPLAY"))
	log.Printf("XAUTHORITY: %s\n", os.Getenv("XAUTHORITY"))
	log.Printf("DBUS_SESSION_BUS_ADDRESS: %s\n", os.Getenv("DBUS_SESSION_BUS_ADDRESS"))
	log.Printf("USER: %s\n", os.Getenv("USER"))
	log.Printf("HOME: %s\n", os.Getenv("HOME"))
	log.Printf("PWD: %s\n", os.Getenv("PWD"))

	fmt.Printf("⏳ Esperando %d segundos a que el sistema esté listo...\n", startDelay)
	log.Printf("Esperando %d segundos...\n", startDelay)
	time.Sleep(time.Duration(startDelay) * time.Second)

	// Verificar DISPLAY
	display := os.Getenv("DISPLAY")
	if display == "" {
		display = ":0"
		os.Setenv("DISPLAY", display)
		log.Println("DISPLAY no configurado, estableciendo a :0")
	}
	fmt.Printf("✓ DISPLAY configurado: %s\n", display)
	log.Printf("DISPLAY: %s\n", display)

	// Buscar Chrome
	chromePath, err := findChrome()
	if err != nil {
		errMsg := fmt.Sprintf("Error: %v", err)
		fmt.Fprintln(os.Stderr, "❌ "+errMsg)
		log.Println("❌ " + errMsg)
		os.Exit(1)
	}
	fmt.Printf("✓ Chrome encontrado: %s\n", chromePath)
	log.Printf("Chrome encontrado: %s\n", chromePath)

	// Argumentos para Chrome
	args := []string{
		fmt.Sprintf("--profile-directory=%s", profileDirectory),
		"--new-window",
		"--no-first-run",
		"--no-default-browser-check",
		unsionURL,
	}

	log.Printf("Argumentos: %v\n", args)
	log.Printf("Comando completo: %s %v\n", chromePath, args)

	// Crear comando
	cmd := exec.Command(chromePath, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	fmt.Printf("🚀 Iniciando Chrome con perfil de Elvia (%s)\n", profileDirectory)
	fmt.Printf("📺 Abriendo: %s\n", unsionURL)
	log.Printf("Ejecutando comando...\n")

	// Iniciar Chrome y ESPERAR a que termine
	if err := cmd.Run(); err != nil {
		errMsg := fmt.Sprintf("Error ejecutando Chrome: %v", err)
		fmt.Fprintln(os.Stderr, "❌ "+errMsg)
		log.Println("❌ " + errMsg)
		os.Exit(1)
	}

	fmt.Println("✅ Chrome se ejecutó")
	log.Println("✅ Chrome terminó su ejecución")
	log.Println("=== Launcher finalizado ===")
}

func findChrome() (string, error) {
	chromeBinaries := []string{
		"google-chrome-stable",
		"google-chrome",
		"chromium-browser",
		"chromium",
	}

	for _, binary := range chromeBinaries {
		if path, err := exec.LookPath(binary); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("no se encontró ninguna instalación de Chrome/Chromium")
}