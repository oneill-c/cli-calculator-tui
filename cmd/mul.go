package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/oneill-c/cli-calculator/internal/compute"
	"github.com/spf13/cobra"
)

var mulCmd = &cobra.Command{
	Use: "mul [...numbers]",
	Short: "Multiply numbers together",
	Args: cobra.MinimumNArgs(2),
	Run: mulFn,
}

func mulFn(cmd *cobra.Command, args []string) {
	nums := []float64{}

	for _, x := range args {
		n, err := strconv.ParseFloat(x, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: invalid number", err)
			os.Exit(2)
		}
		nums = append(nums, n)
	}

	result, err := compute.Mul(nums)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(2)
	}

	fmt.Println(result)
}

func init() {
	rootCmd.AddCommand(mulCmd)
}