// Package cmd contain commands to run as cli
package cmd

import (
	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "cur",
	Short:   "CUR is cloud usage reporting cli",
	Long:    "CUR is cloud usage reporting cli",
	Version: "v0.1.0",
	PreRun: func(cmd *cobra.Command, args []string) {
		xlog.Infoln("hi")
	},
}

// Execute run root command line interface
func Execute() {
	rootCmd.Execute()
}
