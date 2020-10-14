package main

import (
	"github.com/pepodev/cur-cli/cmd"
	"github.com/pepodev/xlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cobra.OnInitialize(initialize)

	cmd.Execute()
}

func initialize() {
	viper.AutomaticEnv()

	viper.AddConfigPath(".")
	viper.SetConfigName("")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		xlog.Debugln(err.Error())
	}
}
