/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/alfaysal/play-with-go/pkg/formatter"
	"github.com/alfaysal/play-with-go/pkg/validation"
	"os"

	"github.com/spf13/cobra"
)

// printerCmd represents the printer command
var printerCmd = &cobra.Command{
	Use:   "printer",
	Short: "A brief description of your command",
	Long:  `Formatter is a CLI tool that helps you manage`,
	Run: func(cmd *cobra.Command, args []string) {
		err, validated := validation.ValidateName("faysal")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(formatter.NameFormatter("All", validated))
	},
}

func init() {
	rootCmd.AddCommand(printerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//printerCmd.PersistentFlags().String("reverse", "", "print ")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
