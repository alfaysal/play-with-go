/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "A brief description of your command",
	Long:  `Form sum.`,
	Args:  cobra.MinimumNArgs(1), // at least user need to provide an arguments
	// Run without error handle
	RunE: sumCommandRunE,
}

func init() {
	arithmeticCmd.AddCommand(sumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func sumCommandRunE(cmd *cobra.Command, args []string) error {

	values := make([]int, len(args))

	for i, v := range args {
		vAsInt, err := strconv.Atoi(v)
		if err != nil { // new code!
			return fmt.Errorf("you provide a value that was not an integer: %s", v)
		}
		values[i] = vAsInt
	}

	showInputs, _ := cmd.Flags().GetBool("show-inputs")

	if showInputs {
		fmt.Fprintln(cmd.OutOrStdout(), "%s\n", strings.Join(args, "+"))
	}

	fmt.Println(getSum(values...))

	return nil
}

func getSum(values ...int) int {
	x := 0
	for _, v := range values {
		x += v
	}

	return x
}
