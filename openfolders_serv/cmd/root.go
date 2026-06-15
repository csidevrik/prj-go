package cmd

import (
	"fmt"
	"os"

	"github.com/adminos/openfolders_serv/internal/storage"
	"github.com/spf13/cobra"
)

var Version = "1.0.0"
var Author = "Carlos Sigua"

var rootCmd = &cobra.Command{
	Use:   "ofolders",
	Short: "📁 CLI tool to manage and open folder groups in Windows Explorer",
	Long: `ofolders - Professional folder management for Windows Explorer

Organize your frequently used folder paths into groups and open them all at once
in new Explorer tabs. Similar to Git workflow but for your files.

Examples:
  ofolders group create work
  ofolders add "D:\\Projects" --alias "Projects" --group "work"
  ofolders run work
`,
	Version: Version,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version and author info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ofolders v%s\n", Version)
		fmt.Printf("Created by %s\n", Author)
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information and statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		groups := store.ListGroups()
		totalPaths := 0

		fmt.Println("📊 ofolders Statistics")
		fmt.Println("════════════════════════════════════")
		fmt.Printf("Total groups: %d\n", len(groups))

		for _, group := range groups {
			totalPaths += len(group.Paths)
		}

		fmt.Printf("Total paths: %d\n", totalPaths)
		fmt.Printf("Config location: %s\n", store.GetConfigPath())
		fmt.Println("════════════════════════════════════")
		fmt.Printf("Created by: %s\n", Author)
		fmt.Printf("Version: v%s\n", Version)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(groupCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(cleanCmd)

	rootCmd.SetVersionTemplate(`ofolders v{{.Version}}
Created by Carlos Sigua
`)
}

func getConfigStore() (*storage.ConfigStore, error) {
	return storage.NewConfigStore()
}
