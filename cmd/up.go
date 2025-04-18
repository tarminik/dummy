/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Ошибка: укажите имя сервиса. Пример: dummy up payment-service")
			return
		}
		service := args[0]
		configPath := "configs/" + service + ".yaml"
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Конфиг %s не найден.\n", configPath)
			return
		}
		fmt.Printf("Запуск окружения для сервиса '%s' (конфиг: %s)...\n", service, configPath)
		// Здесь будет запуск docker compose
		fmt.Println("(Эмуляция запуска: docker compose up -d)")
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
