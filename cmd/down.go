/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop environment for the selected service" ,
	Long: `The down command stops the environment for the selected service (using a config from configs).

Example usage:
  dummy down payment-service
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: please specify a service name. Example: dummy down payment-service")
			return
		}
		service := args[0]
		configPath := "configs/" + service + ".yaml"
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Config %s not found.\n", configPath)
			return
		}
		fmt.Printf("Stopping environment for service '%s' (config: %s)...\n", service, configPath)
		// Here will be the real docker compose down
		fmt.Println("(Simulation: docker compose down)")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
