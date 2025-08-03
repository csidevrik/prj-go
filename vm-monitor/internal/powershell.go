// powershell.go - Ejecuta comandos PowerShell desde Go usando sesión dinámica
package internal

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
