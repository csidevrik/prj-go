package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateGroup string

var updateCmd = &cobra.Command{
	Use:   "update <alias> <new-path>",
	Short: "Update a path in a group",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		alias := args[0]
		newPath := args[1]

		if updateGroup == "" {
			return fmt.Errorf("--group flag is required")
		}

		if err := store.UpdatePath(updateGroup, alias, newPath); err != nil {
			return err
		}

		fmt.Printf("✅ Path updated in group '%s'\n", updateGroup)
		fmt.Printf("   Alias: %s\n", alias)
		fmt.Printf("   New path: %s\n", newPath)
		return nil
	},
}

func init() {
	updateCmd.Flags().StringVar(&updateGroup, "group", "", "Group name (required)")
}
