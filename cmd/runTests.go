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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Получаем значения флагов
		service, _ := cmd.Flags().GetString("service")
		command, _ := cmd.Flags().GetString("command")

		// Проверяем, что оба флага заданы
		if service == "" || command == "" {
			fmt.Println("Ошибка: укажите оба флага: --service и --command. Пример: dummy run-tests --service=payment-service --command=\"pytest tests/\"")
			return
		}
		configPath := "configs/" + service + ".yaml"
		// Проверяем, что конфиг существует
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Конфиг %s не найден.\n", configPath)
			return
		}
		// Эмулируем запуск тестов
		fmt.Printf("Запуск тестов для сервиса '%s' с командой: %s\n", service, command)
		fmt.Println("(Эмуляция: выполнение команды в окружении)")
		fmt.Println("Тесты завершены: PASSED (эмуляция)")
	},
}

func init() {
	rootCmd.AddCommand(runTestsCmd)

	// Регистрируем флаги для команды run-tests
	runTestsCmd.Flags().String("service", "", "Имя сервиса (обязательный флаг)")
	runTestsCmd.Flags().String("command", "", "Команда для запуска тестов (обязательный флаг)")
}
