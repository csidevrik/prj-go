/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getformatMACCmd represents the getformatMAC command
var getformatMACCmd = &cobra.Command{
	Use:   "getformatMAC",
	Short: "Get the format of a MAC address.",
	Long:  `This command returns the format of a given MAC address. The output will be`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getformatMAC called")
	},
}

func init() {
	rootCmd.AddCommand(getformatMACCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getformatMACCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getformatMACCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
