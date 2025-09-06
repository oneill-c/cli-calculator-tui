package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/oneill-c/cli-calculator/internal/compute"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add [...numbers]",
	Short: "Add numbers together",
	Args: cobra.MinimumNArgs(2),
	Run: addFn,
}

func addFn(cmd *cobra.Command, args []string) {
	nums := []float64{}

	for _, x := range args {
		n, err := strconv.ParseFloat(x, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: invalid number", x)
			os.Exit(2)
		}
		nums = append(nums, n)
	}

	result, err := compute.Add(nums)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(2)
	}

	fmt.Println(result)
}

func init() {
    rootCmd.AddCommand(addCmd)
}