/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// saluteCmd represents the salute command
var saluteCmd = &cobra.Command{
	Use:   "salute",
	Short: "Manda un saludo raro",
	Long:  `En realidad es solo un saludo no debes estremcerte antes`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("salute sin nada, solo es un saludo")
	},
}

func init() {
	rootCmd.AddCommand(saluteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saluteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saluteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
