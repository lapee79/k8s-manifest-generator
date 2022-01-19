package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var file string

var rootCmd = &cobra.Command{
	Use:          "k8s-manifest-generator",
	SilenceUsage: true,
	Short:        "Generates the Kustomize manifests.",
	Long: `k8s-manifest-generator generates the Kustomize manifests using 
the application definition JSON file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {}
