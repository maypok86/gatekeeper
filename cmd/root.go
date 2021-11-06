package cmd

import (
	"github.com/spf13/cobra"
)

var serverURL string

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverURL, "url", "u", "localhost:50051", "Gatekeeper server url")
}

var rootCmd = &cobra.Command{
	Use:   "gk",
	Short: "Gatekeeper service and cli",
	Long:  `Gatekeeper is storage limiting service and cli tool`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
