package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "tldwizard",
	Short:   "Changing the TLD for a given domain name",
	Long:    `Swapping out the TLD, you can modify the domain name's extension to better suit your needs.`,
	Version: version,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.AddCommand()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
