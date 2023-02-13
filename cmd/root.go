// Package cmd 程序所有命令
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd 主体命令定义
var rootCmd = &cobra.Command{
	Use:              "snail",
	Short:            "Snail program",
	Long:             "Snail program",
	PersistentPreRun: persistentPreRun,
	Run:              run,
}

// preRun 加载配置文件初始化全局变量, 子命令会继承
func persistentPreRun(cmd *cobra.Command, args []string) {
}

// run 程序执行主体
func run(cmd *cobra.Command, args []string) {

}

// Execute 程序入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
