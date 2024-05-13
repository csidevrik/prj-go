/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rcloneCmd represents the rclone command
var rcloneCmd = &cobra.Command{
	Use:   "rclone",
	Short: "mount remote onedrive disk or google drive disk",
	Long: `rclone mount a remote onedrive or google drive disk. For example:

rclone is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rclone called")
	},
}

func init() {
	rootCmd.AddCommand(rcloneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rcloneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rcloneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
