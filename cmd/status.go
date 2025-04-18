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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configsDir := "configs"
		if len(args) >= 1 {
			service := args[0]
			configPath := configsDir + "/" + service + ".yaml"
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				fmt.Printf("Конфиг %s не найден.\n", configPath)
				return
			}
			fmt.Printf("Сервис %s: Запущен (эмуляция)\n", service)
			return
		}
		files, err := os.ReadDir(configsDir)
		if err != nil {
			fmt.Printf("Ошибка чтения папки %s: %v\n", configsDir, err)
			return
		}
		fmt.Println("Статус всех сервисов:")
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
				name := strings.TrimSuffix(file.Name(), ".yaml")
				fmt.Printf("- %s: Запущен (эмуляция)\n", name)
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
