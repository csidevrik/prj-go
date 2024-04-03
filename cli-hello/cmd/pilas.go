/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pilasCmd represents the pilas command
var pilasCmd = &cobra.Command{
	Use:   "pilas",
	Short: "Te dice pilas en la terminal",
	Long: `El comando imprime en terminal la palabra pilas. For example:
	cli-hello pilas .
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pilas pana")
	},
}

func init() {
	rootCmd.AddCommand(pilasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pilasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pilasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
