/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

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
	arithmeticCmd.AddCommand(subtractCmd)
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

	showInputs, _ := cmd.Flags().GetBool("show-inputs")

	if showInputs {
		fmt.Fprintln(cmd.OutOrStdout(), "%s\n", strings.Join(args, "-"))
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
