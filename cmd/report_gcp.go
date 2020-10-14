package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	reportCmd.AddCommand(gcpCmd)

	gcpCmd.Flags().StringP("region", "", "", "Target region to collect usage")
}

var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Collect cloud usage data and make report",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
