package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "metracker",
	Short: "Metracker is a smart proxy for databases.",
	Long:  `A low-level security tool that intercepts TCP traffic to apply rate limiting and block suspicious queries.`,
}

func executeRootCmd() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	executeRootCmd()
}
