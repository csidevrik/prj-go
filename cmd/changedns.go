package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// changednsCmd representa el comando "changedns"
var changednsCmd = &cobra.Command{
	Use:   "changeDNS1 [IP]",
	Short: "Cambia el DNS primario de la interfaz activa",
	Args:  cobra.ExactArgs(1), // Solo un argumento: la IP del nuevo DNS
	Run: func(cmd *cobra.Command, args []string) {
		newDNS := args[0]
		if !isValidIP(newDNS) {
			fmt.Println("IP inválida:", newDNS)
			return
		}

		// Ejecutamos netsh para cambiar DNS
		fmt.Println("Cambiando DNS primario a:", newDNS)

		// Primero obtener nombre de la interfaz activa
		ifaceName, err := getActiveInterface()
		if err != nil {
			fmt.Println("Error obteniendo interfaz activa:", err)
			return
		}

		// Cambiar el DNS
		command := exec.Command("netsh", "interface", "ip", "set", "dns", "name="+ifaceName, "source=static", "addr="+newDNS)
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error cambiando DNS:", err)
			fmt.Println(string(output))
			return
		}

		fmt.Println("DNS cambiado correctamente.")
	},
}

func init() {
	rootCmd.AddCommand(changednsCmd)
}

// Función auxiliar para validar una IP simple (puedes hacerla más estricta si quieres)
func isValidIP(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || len(p) > 3 {
			return false
		}
	}
	return true
}

// Función para obtener la interfaz activa
func getActiveInterface() (string, error) {
	// Aquí simplificamos usando un comando "route print" para buscar la ruta por defecto
	out, err := exec.Command("powershell", "-Command", "Get-NetRoute -DestinationPrefix 0.0.0.0/0 | Sort-Object -Property RouteMetric | Select-Object -First 1 | Get-NetIPInterface | Select-Object -ExpandProperty InterfaceAlias").Output()
	if err != nil {
		return "", err
	}
	iface := strings.TrimSpace(string(out))
	return iface, nil
}
