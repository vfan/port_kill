package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portkill [port]",
	Short: "端口杀手：根据端口号杀掉对应进程",
	Long:  "一个简单的命令行工具，用于查找并杀掉指定端口的监听进程。",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		port := args[0]
		killByPort(port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
