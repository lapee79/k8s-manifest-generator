package cmd

import (
	"github.com/lapee79/k8s-manifest-generator/generator"
	"github.com/spf13/cobra"
	"log"
)

// generateCmd represents the generate command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Generates the Kustomize manifests.",
	Long: `The generate command generates Kustomize manifests.

Examples:
  k8s-manifest-generator -f app.json`,
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.Run(file)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the Kustomize manifests.",
	Long: `The generate command generates Kustomize manifests.

Examples:
  k8s-manifest-generator -f app.json`,
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.Run(file)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&file, "file", "f", "", "The application definition JSON file")
	err := generateCmd.MarkFlagRequired("file")
	if err != nil {
		return
	}
}
