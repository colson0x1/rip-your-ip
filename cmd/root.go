package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "RIPyourIP - CLI",
		Long:  `RIPyourIP - CLI`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
