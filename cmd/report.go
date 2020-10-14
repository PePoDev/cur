package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reportCmd)
	rootCmd.PersistentFlags().StringP("output", "o", "json", "Output format to make report [json, yaml, slack, telegram]")
	rootCmd.PersistentFlags().IntP("interval", "i", 10, "Schedule to report (minuts)")
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Collect cloud usage data and make report",
}
