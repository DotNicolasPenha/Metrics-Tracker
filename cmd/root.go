package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "metracker",
	Short: "Metracker is a high-performance database proxy.",
	Long:  "A low-level security tool that intercepts TCP traffic to apply rate limiting and block suspicious queries.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
