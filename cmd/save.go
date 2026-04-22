package cmd

import (
	"fmt"
	"os"

	"github.com/DotNicolasPenha/Metrics-Tracker/interceptor"
	"github.com/DotNicolasPenha/Metrics-Tracker/user"
	"github.com/spf13/cobra"
)

var (
	configFilePath string
	maxConn        int
	blockQuery     string
	retrys         int
	proxyAddr      string
	dbAddr         string
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save or update configurations via flags.",
	Run: func(cmd *cobra.Command, args []string) {
		var currentUser user.User

		if _, err := os.Stat(configFilePath); err == nil {
			loadedUser, err := user.LoadUser(configFilePath)
			if err == nil {
				currentUser = loadedUser
			}
		}

		if cmd.Flags().Changed("max-conn") {
			currentUser.Configurations.Limits.MaxActConnections = maxConn
		}

		if cmd.Flags().Changed("block") {
			newRule := user.BlockQuerie{
				Query:  []byte(blockQuery),
				Retrys: retrys,
			}
			currentUser.Configurations.BlockQueries = append(currentUser.Configurations.BlockQueries, newRule)
			fmt.Printf("Added block rule: %s (Retries: %d)\n", blockQuery, retrys)
		}

		if cmd.Flags().Changed("proxy-addr") && cmd.Flags().Changed("db-addr") {
			newInterceptor := &interceptor.Interceptor{
				ProxyAddr: proxyAddr,
				DBAddr:    dbAddr,
			}
			currentUser.Interceptors = append(currentUser.Interceptors, newInterceptor)
			fmt.Printf("Added new interceptor: %s -> %s\n", proxyAddr, dbAddr)
		} else if cmd.Flags().Changed("proxy-addr") || cmd.Flags().Changed("db-addr") {
			fmt.Println("Warning: To add an interceptor, you must provide both --proxy-addr and --db-addr")
		}

		if currentUser.Interceptors == nil {
			currentUser.Interceptors = []*interceptor.Interceptor{}
		}

		err := user.SaveUser(currentUser, configFilePath)
		if err != nil {
			fmt.Printf("Error saving configuration: %v\n", err)
			return
		}

		fmt.Printf("Configuration successfully updated in: %s\n", configFilePath)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringVarP(&configFilePath, "file", "f", "config.json", "Target configuration file")
	saveCmd.Flags().IntVarP(&maxConn, "max-conn", "m", 100, "Update maximum active connections")

	saveCmd.Flags().StringVarP(&blockQuery, "block", "b", "", "Add a new query string to blacklist")
	saveCmd.Flags().IntVarP(&retrys, "retrys", "r", 3, "Number of retries allowed for the blocked query before permanent ban")

	saveCmd.Flags().StringVar(&proxyAddr, "proxy-addr", "pa", "Local address for the proxy to listen on (e.g., :5433)")
	saveCmd.Flags().StringVar(&dbAddr, "db-addr", "da", "Target database address (e.g., localhost:5432)")
}
