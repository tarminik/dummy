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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the list of all available services (configs)",
	Long: `The list command shows all services that have configs in the configs folder.

Example usage:
  dummy list
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Получаем список файлов из папки configs
		configsDir := "configs"
		files, err := os.ReadDir(configsDir)
		if err != nil {
			fmt.Printf("Error reading directory %s: %v\n", configsDir, err)
			return
		}
		fmt.Println("Available configurations:")
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".yaml") {
				name := strings.TrimSuffix(file.Name(), ".yaml")
				fmt.Println("-", name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
