/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tarminik/dummy/internal/config"
	"github.com/tarminik/dummy/internal/compose"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop environment for the selected service" ,
	Long: `The down command stops the environment for the selected service (using a config from configs).

Example usage:
  dummy down payment-service
`,
	// Command handler for 'down'. Checks args, config existence, and simulates stop.
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: please specify a service name. Example: dummy down payment-service")
			return
		}
		service := args[0]
		configsDir := "configs"
		if !config.ConfigExists(configsDir, service) {
			fmt.Printf("Config %s/%s.yaml not found.\n", configsDir, service)
			return
		}
		// Use the compose package to simulate stopping the service
		compose.Down(service, configsDir+"/"+service+".yaml")
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
