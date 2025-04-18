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

// runTestsCmd represents the runTests command
var runTestsCmd = &cobra.Command{
	Use:   "run-tests",
	Short: "Run tests in the selected service environment" ,
	Long: `The run-tests command runs tests inside the environment of the selected service.

Example usage:
  dummy run-tests --service=payment-service --command="pytest tests/"
`,
	// Command handler for 'run-tests'. Checks flags, config, and simulates tests.
	Run: func(cmd *cobra.Command, args []string) {
		// Get flag values
		service, _ := cmd.Flags().GetString("service")
		command, _ := cmd.Flags().GetString("command")

		// Check that both flags are set
		if service == "" || command == "" {
			fmt.Println("Error: please specify both flags: --service and --command. Example: dummy run-tests --service=payment-service --command=\"pytest tests/\"")
			return
		}
		configsDir := "configs"
		if !config.ConfigExists(configsDir, service) {
			fmt.Printf("Config %s/%s.yaml not found.\n", configsDir, service)
			return
		}
		// Use the compose package to simulate running tests
		compose.RunTests(service, command)
	},
}

func init() {
	rootCmd.AddCommand(runTestsCmd)

	// Регистрируем флаги для команды run-tests
	runTestsCmd.Flags().String("service", "", "Имя сервиса (обязательный флаг)")
	runTestsCmd.Flags().String("command", "", "Команда для запуска тестов (обязательный флаг)")
}
