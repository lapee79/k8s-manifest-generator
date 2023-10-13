package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "v0.1.2"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long: `Print the version information.

Examples:
  k8s-manifest-generator version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
