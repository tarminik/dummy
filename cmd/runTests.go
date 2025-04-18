/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// runTestsCmd represents the runTests command
var runTestsCmd = &cobra.Command{
	Use:   "run-tests",
	Short: "Run tests in the selected service environment" ,
	Long: `The run-tests command runs tests inside the environment of the selected service.

Example usage:
  dummy run-tests --service=payment-service --command="pytest tests/"
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Получаем значения флагов
		service, _ := cmd.Flags().GetString("service")
		command, _ := cmd.Flags().GetString("command")

		// Проверяем, что оба флага заданы
		if service == "" || command == "" {
			fmt.Println("Error: please specify both flags: --service and --command. Example: dummy run-tests --service=payment-service --command=\"pytest tests/\"")
			return
		}
		configPath := "configs/" + service + ".yaml"
		// Проверяем, что конфиг существует
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Config %s not found.\n", configPath)
			return
		}
		// Эмулируем запуск тестов
		fmt.Printf("Running tests for service '%s' with command: %s\n", service, command)
		fmt.Println("(Simulation: running command in environment)")
		fmt.Println("Tests finished: PASSED (simulation)")
	},
}

func init() {
	rootCmd.AddCommand(runTestsCmd)

	// Регистрируем флаги для команды run-tests
	runTestsCmd.Flags().String("service", "", "Имя сервиса (обязательный флаг)")
	runTestsCmd.Flags().String("command", "", "Команда для запуска тестов (обязательный флаг)")
}
