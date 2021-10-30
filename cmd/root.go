package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gk",
	Short: "Gatekeeper service and cli",
	Long:  `Gatekeeper is rate limiting service and cli tool`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
