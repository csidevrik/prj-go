package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addGroup string
	addAlias string
)

var addCmd = &cobra.Command{
	Use:   "add <path>",
	Short: "Add a path to a group",
	Long:  `Add a folder path to an existing group with an optional alias`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		path := args[0]

		if addGroup == "" {
			return fmt.Errorf("--group flag is required")
		}

		if addAlias == "" {
			return fmt.Errorf("--alias flag is required")
		}

		if err := store.AddPath(addGroup, path, addAlias); err != nil {
			return err
		}

		fmt.Printf("✅ Path added to group '%s'\n", addGroup)
		fmt.Printf("   Alias: %s\n", addAlias)
		fmt.Printf("   Path: %s\n", path)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVar(&addGroup, "group", "", "Group name (required)")
	addCmd.Flags().StringVar(&addAlias, "alias", "", "Path alias (required)")
}
