package cmd

import (
	"bookkeeping-shell/funcs"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(viewCmd)
}

// 输出指定格式的文件
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "展示图表",
	Long:  "在浏览器中展示图表",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := funcs.GenViewJS();
		if err != nil {
			fmt.Printf("error: %v:\n", errors.Unwrap(err))
		}
		fmt.Println("正在打开浏览器 ...")
	},
}
