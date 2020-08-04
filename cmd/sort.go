package cmd

import (
"bookkeeping-shell/funcs"
"errors"
"fmt"
"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(sortCmd)
}

// 输出指定格式的文件
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "排序",
	Long:  "对 text 文件数据进行排序",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := funcs.Sort();
		if err != nil {
			fmt.Printf("error: %v:\n", errors.Unwrap(err))
		}
	},
}
