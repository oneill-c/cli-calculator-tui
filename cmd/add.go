package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add [...numbers]",
	Short: "Add numbers together",
	Args: cobra.MinimumNArgs(2),
	Run: addFn,
}

func addFn(cmd *cobra.Command, args []string) {
	fmt.Println("Hello World from addFn") 
}

func init() {
    rootCmd.AddCommand(addCmd)
}