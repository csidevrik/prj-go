package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeGroup string

var removeCmd = &cobra.Command{
	Use:   "remove <alias>",
	Short: "Remove a path from a group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		alias := args[0]

		if removeGroup == "" {
			return fmt.Errorf("--group flag is required")
		}

		if err := store.RemovePath(removeGroup, alias); err != nil {
			return err
		}

		fmt.Printf("✅ Path with alias '%s' removed from group '%s'\n", alias, removeGroup)
		return nil
	},
}

func init() {
	removeCmd.Flags().StringVar(&removeGroup, "group", "", "Group name (required)")
}
