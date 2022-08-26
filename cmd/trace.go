package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP.",
	Long:  `Trace the IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
			}
		} else {
			fmt.Println("Provide an IP address ayy!")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
