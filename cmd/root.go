package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "calc",
    Short: "A simple CLI calculator",
    Long:  "CLI calculator with subcommands and optional TUI.",
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
	}
}