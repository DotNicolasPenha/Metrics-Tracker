package cmd

import (
	"fmt"
	"os"

	"github.com/DotNicolasPenha/Metrics-Tracker/interceptor"
	"github.com/DotNicolasPenha/Metrics-Tracker/user"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the specified interceptor.",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(configFilePath); err != nil {
			fmt.Printf("Error: Config file not found at '%s'. Please create it using 'save' command.\n", configFilePath)
			return
		}

		currentUser, err := user.LoadUser(configFilePath)
		if err != nil {
			fmt.Printf("Error loading config file: %v\n", err)
			return
		}

		var target *interceptor.Interceptor
		for _, i := range currentUser.Interceptors {
			if i.Name == name {
				target = i
				break
			}
		}

		if target == nil {
			fmt.Printf("Error: Interceptor '%s' not found in config. Please create it using 'save' command.\n", name)
			return
		}

		fmt.Printf("Running interceptor '%s' on %s -> %s\n", target.Name, target.ProxyAddr, target.DBAddr)
		target.Run()
	},
}

func init() {
	runCmd.Flags().StringVarP(&configFilePath, "config", "c", "config.json", "Path to the configuration file")
	runCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the interceptor to run")
	runCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(runCmd)
}
