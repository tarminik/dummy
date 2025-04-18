package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Version of the dummy CLI. Update this value for each release.
const cliVersion = "0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version of the dummy CLI tool",
	Long: `Prints the version of the dummy CLI tool.
Useful for bug reports and checking updates.

Example usage:
  dummy version
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("dummy version %s\n", cliVersion)
	},
}

func init() {
	// Register the version command with the root command.
	rootCmd.AddCommand(versionCmd)
}
