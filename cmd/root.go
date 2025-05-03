package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd es el comando principal de tu CLI.
var RootCmd = &cobra.Command{
	Use:   "perfilizer",
	Short: "Herramienta CLI para gestionar tu $PROFILE.ps1",
	Run: func(cmd *cobra.Command, args []string) {
		// Si ejecutas `perfilizer` sin subcomando,
		// muestra la ayuda:
		cmd.Help()
	},
}

// Execute arranca la ejecución de RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// init registra los subcomandos de RootCmd.
func init() {
	// subcomando de completions
	RootCmd.AddCommand(completionCmd)
}

// completionCmd genera el script de completions para distintos shells.
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Genera el script de completions para el shell especificado",
	Long: `Para PowerShell, puedes hacer:
  perfilizer completion powershell > perfilizer-completion.ps1
  # Luego, añade al final de tu $PROFILE:
  . \"<ruta>\\perfilizer-completion.ps1\"

Cada vez que abras PowerShell tendrás autocompletado de perfilizer.
`,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return RootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return RootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return RootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			// GenPowerShellCompletionWithDesc incluye descripciones en el script
			return RootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		}
		return fmt.Errorf("shell %q no soportado", args[0])
	},
}
