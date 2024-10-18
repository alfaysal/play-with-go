/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// birthdayCmd represents the birthday command
var birthdayCmd = &cobra.Command{
	Use:   "birthday",
	Short: "A short for birthday",
	Long:  `birthday`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("birthday called")
	},
}

func init() {
	rootCmd.AddCommand(birthdayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// birthdayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// birthdayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
