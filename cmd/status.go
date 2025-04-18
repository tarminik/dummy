/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of all services or a selected service" ,
	Long: `The status command shows the status of all services (by default) or a specific service (if given).

Examples:
  dummy status
  dummy status payment-service
`,
	Run: func(cmd *cobra.Command, args []string) {
		configsDir := "configs"
		if len(args) >= 1 {
			service := args[0]
			configPath := configsDir + "/" + service + ".yaml"
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				fmt.Printf("Config %s not found.\n", configPath)
				return
			}
			fmt.Printf("Service %s: Running (simulation)\n", service)
			return
		}
		files, err := os.ReadDir(configsDir)
		if err != nil {
			fmt.Printf("Error reading directory %s: %v\n", configsDir, err)
			return
		}
		fmt.Println("Status of all services:")
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
				name := strings.TrimSuffix(file.Name(), ".yaml")
				fmt.Printf("- %s: Running (simulation)\n", name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
