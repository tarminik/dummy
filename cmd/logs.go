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

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show logs for the selected service" ,
	Long: `The logs command shows logs for the selected service (by config from configs).

Example usage:
  dummy logs payment-service
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: please specify a service name. Example: dummy logs payment-service")
			return
		}
		service := args[0]
		configsDir := "configs"
		if !config.ConfigExists(configsDir, service) {
			fmt.Printf("Config %s/%s.yaml not found.\n", configsDir, service)
			return
		}
		// Use the compose package to simulate logs output
		compose.Logs(service)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
