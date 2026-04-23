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
	name           string
	maxConn        int
	blockQuery     string
	retrys         int
	proxyAddr      string
	dbAddr         string
	authorizedIP   string
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save or update configurations per interceptor.",
	Run: func(cmd *cobra.Command, args []string) {
		var currentUser user.User

		if _, err := os.Stat(configFilePath); err == nil {
			if loadedUser, err := user.LoadUser(configFilePath); err == nil {
				currentUser = loadedUser
			}
		}

		var target *interceptor.Interceptor
		for _, i := range currentUser.Interceptors {
			if i.Name == name {
				target = i
				break
			}
		}

		if target == nil && cmd.Flags().Changed("proxy-addr") && cmd.Flags().Changed("db-addr") {
			target = &interceptor.Interceptor{
				Name:      name,
				ProxyAddr: proxyAddr,
				DBAddr:    dbAddr,
			}
			currentUser.Interceptors = append(currentUser.Interceptors, target)
			fmt.Printf("Created new interceptor: %s\n", name)
		}

		if target == nil {
			fmt.Printf("Error: Interceptor '%s' not found. Provide --proxy-addr and --db-addr to create it.\n", name)
			return
		}

		if cmd.Flags().Changed("max-conn") {
			target.Configurations.Limits.MaxActConnections = maxConn
		}

		if cmd.Flags().Changed("block") {
			newRule := interceptor.BlockQuerie{
				Query:  []byte(blockQuery),
				Retrys: retrys,
			}
			target.Configurations.BlockQueries = append(target.Configurations.BlockQueries, newRule)
		}

		if cmd.Flags().Changed("authorized-ips") {
			target.Configurations.AuthorizedIPs = append(target.Configurations.AuthorizedIPs, authorizedIP)
		}

		if err := user.SaveUser(currentUser, configFilePath); err != nil {
			fmt.Printf("Error saving configuration: %v\n", err)
			return
		}

		fmt.Printf("Interceptor '%s' successfully updated in: %s\n", name, configFilePath)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringVarP(&configFilePath, "file", "f", "config.json", "Configuration file")
	saveCmd.Flags().StringVarP(&name, "name", "n", "default", "Name of the interceptor to manage")

	saveCmd.Flags().IntVarP(&maxConn, "max-conn", "m", 100, "Max connections for this interceptor")
	saveCmd.Flags().StringVarP(&blockQuery, "block", "b", "", "Add query to blacklist")
	saveCmd.Flags().IntVarP(&retrys, "retrys", "r", 3, "Retry limit for the blocked query")
	saveCmd.Flags().StringVarP(&authorizedIP, "authorized-ips", "i", "", "Comma-separated list of authorized IPs")

	saveCmd.Flags().StringVar(&proxyAddr, "proxy-addr", "", "Local proxy address")
	saveCmd.Flags().StringVar(&dbAddr, "db-addr", "", "Target database address")
}
