/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chaoCmd represents the chao command
var chaoCmd = &cobra.Command{
	Use:   "chao",
	Short: "es un comando que te dice adios en pantalla",
	Long: `Es un comando para decir adios en terminal. For example:

Adios pana.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chao pana")
	},
}

func init() {
	rootCmd.AddCommand(chaoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chaoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chaoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
