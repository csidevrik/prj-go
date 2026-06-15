package cmd

import (
	"fmt"

	"github.com/adminos/openfolders_serv/internal/explorer"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <group>",
	Short: "Open a group in Windows Explorer tabs",
	Long:  `Open all paths from a group as tabs in a new Explorer window`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		groupName := args[0]
		group, err := store.GetGroup(groupName)
		if err != nil {
			return err
		}

		if len(group.Paths) == 0 {
			return fmt.Errorf("group '%s' has no paths", groupName)
		}

		paths := make([]string, len(group.Paths))
		for i, p := range group.Paths {
			paths[i] = p.Path
		}

		fmt.Printf("🚀 Opening group '%s' with %d tabs...\n", groupName, len(paths))
		for i, path := range paths {
			fmt.Printf("   [%d] %s\n", i+1, path)
		}

		if err := explorer.OpenExplorerWithPaths(paths); err != nil {
			return fmt.Errorf("failed to open explorer: %w", err)
		}

		fmt.Println("✅ Done!")
		return nil
	},
}
