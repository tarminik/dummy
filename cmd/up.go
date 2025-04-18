/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"dummy/internal/config"
	"dummy/internal/compose"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start environment for the selected service (by config)",
	Long: `The up command starts the environment for the selected service using a yaml config from the configs folder.

Example usage:
  dummy up payment-service
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: please specify a service name. Example: dummy up payment-service")
			return
		}
		service := args[0]
		configsDir := "configs"
		if !config.ConfigExists(configsDir, service) {
			fmt.Printf("Config %s/%s.yaml not found.\n", configsDir, service)
			return
		}
		// Use the compose package to simulate starting the service
		compose.Up(service, configsDir+"/"+service+".yaml")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
