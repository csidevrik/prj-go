/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/carlossiguam/prj-go/cli-g/lib"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "lista los formatos de mac address encontrados en formats.json",
	Long:  `Show a list of formats exists`,
	Run: func(cmd *cobra.Command, args []string) {
		formatos, err := lib.ObtenerFormatos()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Puedes imprimir o manipular la lista de formatos aquí
		fmt.Println("Formatos:")
		for _, formato := range formatos {
			fmt.Printf("- %s: %s\n", formato.NAME, formato.MACFORMAT)
		}
		// fmt.Println("uno ")
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
