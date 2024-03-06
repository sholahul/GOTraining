package cmd

import (
	"os"

	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "app",
	PreRun: func(cmd *cobra.Command, args []string) {
		doMigration()
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error()
		os.Exit(0)
	}
}
