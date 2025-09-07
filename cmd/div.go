package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/oneill-c/cli-calculator/internal/compute"
	"github.com/spf13/cobra"
)

var divCmd = &cobra.Command{
	Use: "div <a> <b>",
	Short: "Divide two numbers",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("div requires exactly 2 numbers, got %d", len(args))
		}
		return nil
	},
	Run: divFn,
}

func divFn(cmd *cobra.Command, args []string) {
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid number")
		os.Exit(2)
	}

	b, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid number")
		os.Exit(2)
	}

	result, err := compute.Div(a, b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(2)
	}

	fmt.Println(result)
}

func init() {
	rootCmd.AddCommand(divCmd)
}