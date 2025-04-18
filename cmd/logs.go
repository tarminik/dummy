/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
		// Проверяем, что передан сервис
		if len(args) < 1 {
			fmt.Println("Error: please specify a service name. Example: dummy logs payment-service")
			return
		}
		service := args[0]
		configPath := "configs/" + service + ".yaml"
		// Проверяем, что конфиг существует
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Config %s not found.\n", configPath)
			return
		}
		// Эмулируем вывод логов
		fmt.Printf("Logs for service '%s':\n", service)
		fmt.Println("[INFO] Service started successfully.")
		fmt.Println("[INFO] Handling request on /api/v1/ping")
		fmt.Println("[ERROR] Database connection timeout (simulation)")
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
