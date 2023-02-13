package cmd

import (
	"github.com/daymenu/snail/pkg/console"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// versionCmd 版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Snail",
	Long:  "Print the version number of Snail",
	Run: func(cmd *cobra.Command, args []string) {
		console.Success("Snail v1.0")
	},
}
