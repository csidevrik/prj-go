package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage folder groups",
	Long:  `Create, delete, and list folder groups`,
}

var groupCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		groupName := args[0]
		if err := store.CreateGroup(groupName); err != nil {
			return err
		}

		fmt.Printf("✅ Group '%s' created successfully\n", groupName)
		return nil
	},
}

var groupDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		groupName := args[0]
		if err := store.DeleteGroup(groupName); err != nil {
			return err
		}

		fmt.Printf("✅ Group '%s' deleted successfully\n", groupName)
		return nil
	},
}

var groupListCmd = &cobra.Command{
	Use:   "list [name]",
	Short: "List groups or paths in a group",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := getConfigStore()
		if err != nil {
			return err
		}

		if len(args) == 0 {
			// List all groups
			groups := store.ListGroups()
			if len(groups) == 0 {
				fmt.Println("📭 No groups found")
				return nil
			}

			fmt.Println("📁 Available Groups:")
			fmt.Println("════════════════════════════════════")
			for name, group := range groups {
				fmt.Printf("  • %s (%d paths)\n", name, len(group.Paths))
			}
			return nil
		}

		// List paths in specific group
		groupName := args[0]
		group, err := store.GetGroup(groupName)
		if err != nil {
			return err
		}

		fmt.Printf("📂 Group: %s\n", groupName)
		fmt.Println("════════════════════════════════════")

		if len(group.Paths) == 0 {
			fmt.Println("  (empty)")
			return nil
		}

		for i, path := range group.Paths {
			fmt.Printf("  %d. [%s] %s\n", i+1, path.Alias, path.Path)
		}
		return nil
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)
	groupCmd.AddCommand(groupDeleteCmd)
	groupCmd.AddCommand(groupListCmd)
}
