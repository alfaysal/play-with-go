/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "A brief description of your command",
	Long:  `Subtract`,
	RunE:  subCommandRunE,
}

func init() {
	rootCmd.AddCommand(subtractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	subtractCmd.Flags().BoolP("invert-sign", "i", false, "inverts the sign of the result.")
}

func subCommandRunE(cmd *cobra.Command, args []string) error {

	values := make([]int, len(args))

	for i, v := range args {
		vAsInt, err := strconv.Atoi(v)
		if err != nil { // new code!
			return fmt.Errorf("you provide a value that was not an integer: %s", v)
		}
		values[i] = vAsInt
	}

	// get the invert sign flag
	result := getSub(values...)
	invert, _ := cmd.Flags().GetBool("invert-sign")

	if invert {
		result = -result
	}

	fmt.Println(result)
	return nil
}

func getSub(values ...int) int {
	fmt.Println(values)
	x := 0
	for _, v := range values {
		x -= v
	}

	return x
}
