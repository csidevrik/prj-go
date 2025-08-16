// powershell.go - Ejecuta comandos PowerShell desde Go usando sesión dinámica
package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecutePowerShellWithAuth(server, user, password, vmCommand string) string {
	tempScript := filepath.Join(os.TempDir(), "check_vm_status.ps1")

	fullScript := fmt.Sprintf(`
Set-PowerCLIConfiguration -Scope User -ParticipateInCEIP $false -Confirm:$false | Out-Null
$securePassword = ConvertTo-SecureString "%s" -AsPlainText -Force
$cred = New-Object System.Management.Automation.PSCredential ("%s", $securePassword)
Connect-VIServer -Server "%s" -Credential $cred | Out-Null
%s
`, password, user, server, vmCommand)

	err := os.WriteFile(tempScript, []byte(fullScript), 0600)
	if err != nil {
		return fmt.Sprintf("❌ Error al escribir el script PowerShell: %v", err)
	}

	cmd := exec.Command("pwsh", "-File", tempScript)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("❌ Error ejecutando PowerShell: %s\n\n%s", err, string(output))
	}
	return string(output)
}

func GenerateStopVMCommand(vmNames []string) string {
	var quoted []string
	for _, name := range vmNames {
		name = clean(name)
		if name != "" {
			quoted = append(quoted, fmt.Sprintf("\"%s\"", name))
		}
	}
	// return fmt.Sprintf("$vmNames = @(%s)\n$vmNames | ForEach-Object { Stop-VM -VM $_ -Confirm:$false }", strings.Join(quoted, ", "))
	return fmt.Sprintf("$vmNames = @(%s)\n$vmNames | ForEach-Object { Stop-VM -VM $_ -Confirm:$false }", strings.Join(quoted, ", "))

}

func clean(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\r", ""))
}