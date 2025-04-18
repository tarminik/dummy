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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Проверяем, что передан сервис
		if len(args) < 1 {
			fmt.Println("Ошибка: укажите имя сервиса. Пример: dummy logs payment-service")
			return
		}
		service := args[0]
		configPath := "configs/" + service + ".yaml"
		// Проверяем, что конфиг существует
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Конфиг %s не найден.\n", configPath)
			return
		}
		// Эмулируем вывод логов
		fmt.Printf("Логи сервиса '%s':\n", service)
		fmt.Println("[INFO] Service started successfully.")
		fmt.Println("[INFO] Handling request on /api/v1/ping")
		fmt.Println("[ERROR] Database connection timeout (имитация)")
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
