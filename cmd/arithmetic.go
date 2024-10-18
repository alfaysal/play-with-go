/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// arithmeticCmd represents the arithmetic command
var arithmeticCmd = &cobra.Command{
	Use:   "arithmetic",
	Short: "A brief description of your command",
	Long:  `Arithmetic.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("arithmetic called")
	},
}

func init() {
	rootCmd.AddCommand(arithmeticCmd)
	arithmeticCmd.PersistentFlags().BoolP("show-inputs", "s", false, "whether to print inputs")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// arithmeticCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// arithmeticCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
