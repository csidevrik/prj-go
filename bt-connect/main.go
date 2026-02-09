package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	// MAC address del JBL Clip 4
	deviceMAC = "F8:5C:7E:4A:C2:54"
	deviceName = "JBL Clip 4"
	
	// Configuración de reintentos
	maxRetries = 15
	retryDelay = 2 * time.Second
	startDelay = 8 * time.Second // Esperar más que el TV launcher
)

func main() {
	// Configurar logging
	logFile, err := os.OpenFile("/tmp/bt-connect.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("\n=== 🔊 Bluetooth Connector INICIADO [%s] ===\n", timestamp)
	fmt.Println("=== 🔊 Conectando JBL Clip 4 ===")

	// Esperar a que el sistema esté listo
	fmt.Printf("⏳ Esperando %v a que Bluetooth esté listo...\n", startDelay)
	log.Printf("Esperando %v...\n", startDelay)
	time.Sleep(startDelay)

	// Verificar que bluetoothctl esté disponible
	if _, err := exec.LookPath("bluetoothctl"); err != nil {
		errMsg := "bluetoothctl no encontrado"
		fmt.Println("❌ " + errMsg)
		log.Println("❌ " + errMsg)
		os.Exit(1)
	}
	log.Println("✓ bluetoothctl encontrado")

	// Encender Bluetooth
	fmt.Println("🔌 Encendiendo Bluetooth...")
	log.Println("Encendiendo Bluetooth...")
	if err := runBluetoothctl("power on"); err != nil {
		log.Printf("Advertencia al encender Bluetooth: %v\n", err)
	}
	time.Sleep(2 * time.Second)

	// Verificar si el dispositivo ya está conectado
	if isDeviceConnected(deviceMAC) {
		fmt.Printf("✅ %s ya está conectado\n", deviceName)
		log.Printf("✅ Dispositivo %s (%s) ya está conectado\n", deviceName, deviceMAC)
		return
	}

	// Intentar conectar con reintentos
	fmt.Printf("🔍 Buscando %s (%s)...\n", deviceName, deviceMAC)
	log.Printf("Intentando conectar dispositivo %s (%s)\n", deviceName, deviceMAC)

	for attempt := 1; attempt <= maxRetries; attempt++ {
		fmt.Printf("Intento %d/%d...\n", attempt, maxRetries)
		log.Printf("Intento de conexión %d/%d\n", attempt, maxRetries)

		// Intentar conectar
		if err := connectDevice(deviceMAC); err != nil {
			log.Printf("Intento %d falló: %v\n", attempt, err)
			
			if attempt < maxRetries {
				fmt.Printf("⏳ Reintentando en %v...\n", retryDelay)
				time.Sleep(retryDelay)
				continue
			}
			
			// Último intento falló
			errMsg := fmt.Sprintf("No se pudo conectar después de %d intentos", maxRetries)
			fmt.Println("❌ " + errMsg)
			log.Println("❌ " + errMsg)
			os.Exit(1)
		}

		// Conexión exitosa
		fmt.Printf("✅ %s conectado exitosamente!\n", deviceName)
		log.Printf("✅ Dispositivo conectado exitosamente en intento %d\n", attempt)
		log.Println("=== Bluetooth Connector finalizado ===\n")
		return
	}
}

// runBluetoothctl ejecuta un comando en bluetoothctl
func runBluetoothctl(command string) error {
	cmd := exec.Command("bluetoothctl", strings.Split(command, " ")...)
	output, err := cmd.CombinedOutput()
	log.Printf("bluetoothctl %s: %s\n", command, string(output))
	return err
}

// isDeviceConnected verifica si el dispositivo ya está conectado
func isDeviceConnected(mac string) bool {
	cmd := exec.Command("bluetoothctl", "info", mac)
	output, err := cmd.Output()
	
	if err != nil {
		return false
	}

	outputStr := string(output)
	return strings.Contains(outputStr, "Connected: yes")
}

// connectDevice intenta conectar el dispositivo
func connectDevice(mac string) error {
	// Primero, asegurarse de que el dispositivo esté en la lista
	log.Printf("Ejecutando: bluetoothctl connect %s\n", mac)
	
	cmd := exec.Command("bluetoothctl", "connect", mac)
	output, err := cmd.CombinedOutput()
	
	outputStr := string(output)
	log.Printf("Salida de connect: %s\n", outputStr)

	if err != nil {
		// Verificar si aún así se conectó
		time.Sleep(2 * time.Second)
		if isDeviceConnected(mac) {
			return nil
		}
		return fmt.Errorf("error de conexión: %v - %s", err, outputStr)
	}

	// Verificar que realmente se conectó
	time.Sleep(2 * time.Second)
	if !isDeviceConnected(mac) {
		return fmt.Errorf("el comando no reportó error pero el dispositivo no está conectado")
	}

	return nil
}