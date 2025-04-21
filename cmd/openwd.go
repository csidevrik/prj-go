package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// openwdCmd genera y añade una función de PowerShell en el perfil para abrir un directorio.
var openwdCmd = &cobra.Command{
	Use:   "openwd [ruta]",
	Short: "Crea una función en tu $PROFILE para abrir la ruta especificada",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fullPath := args[0]

		// Derivar el nombre de la función a partir de la carpeta final (sin prefijo "prj-")
		base := filepath.Base(fullPath)
		name := strings.TrimPrefix(base, "prj-")
		funcName := "Open" + strings.Title(name)

		// Nombre de la variable interna
		varName := "PATH_" + strings.ToUpper(name)

		// Obtener la ruta al perfil de PowerShell
		// Usa la variable de entorno $PROFILE si está definida
		profilePath := os.Getenv("PROFILE")
		if profilePath == "" {
			usr, err := user.Current()
			if err != nil {
				return err
			}
			profilePath = filepath.Join(usr.HomeDir, "Documents", "PowerShell", "Microsoft.PowerShell_profile.ps1")
		}

		// Determinar ruta relativa si está dentro de $HOME
		relPath := fullPath
		usr, _ := user.Current()
		if strings.HasPrefix(fullPath, usr.HomeDir) {
			rel, err := filepath.Rel(usr.HomeDir, fullPath)
			if err == nil {
				relPath = filepath.ToSlash(rel)
			}
		}

		// Construir el snippet de PowerShell
		snippet := fmt.Sprintf(
			`function %s {
    $%s = Join-Path $env:USERPROFILE "%s"

    if (-not (Test-Path -Path $%s -PathType Container)) {
        New-Item -Path $%s -ItemType Directory
    }

    Set-Location -Path $%s
}
`, funcName, varName, relPath, varName, varName, varName)

		// Abrir (o crear) el archivo de perfil y anexar el snippet
		f, err := os.OpenFile(profilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := f.WriteString("\n" + snippet); err != nil {
			return err
		}

		fmt.Printf("✔ Función %s añadida en %s\n", funcName, profilePath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(openwdCmd)
}
