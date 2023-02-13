// Package console 命令行辅助方法
package console

import (
	"os"

	"github.com/fatih/color"
)

// Success 打印一条成功消息
func Success(msg string) {
	color.Green(msg)
}

// Info 打印一条提示消息
func Info(msg string) {
	color.Cyan(msg)
}

// Warning 打印一条警告消息
func Warning(msg string) {
	color.Yellow(msg)
}

// Error 打印一条报错消息
func Error(msg string) {
	color.Red(msg)
}

// Exit 打印一条错误消息并退出
func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}
