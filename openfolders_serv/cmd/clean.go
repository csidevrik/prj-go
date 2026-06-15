package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanAll bool

var cleanCmd = &cobra.Command{
	Use:   "clean [group]",
	Short: "Clean paths from a group or all groups",
	Long:  `Remove all paths from a group, or remove all paths from all groups if --all is used`,
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		if cleanAll {
			if err := store.CleanAll(); err != nil {
				return err
			}
			fmt.Println("✅ All groups cleaned successfully")
			return nil
		}

		if len(args) == 0 {
			return fmt.Errorf("group name required or use --all flag")
		}

		groupName := args[0]
		if err := store.CleanGroup(groupName); err != nil {
			return err
		}

		fmt.Printf("✅ Group '%s' cleaned successfully\n", groupName)
		return nil
	},
}

func init() {
	cleanCmd.Flags().BoolVar(&cleanAll, "all", false, "Clean all groups")
}
