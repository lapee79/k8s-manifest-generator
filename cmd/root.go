package cmd

import (
	"github.com/lapee79/k8s-manifest-generator/logger"
	"github.com/spf13/cobra"
)

var file string

var rootCmd = &cobra.Command{
	Use:          "k8s-manifest-generator",
	SilenceUsage: true,
	Short:        "Generates the Kustomize manifests.",
	Long: `k8s-manifest-generator generates the Kustomize manifests using 
the application definition JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	logger.Error(err)
}

func init() {}
