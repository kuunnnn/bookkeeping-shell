package cmd

import (
	"bookkeeping-shell/funcs"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var line uint

func init() {
	deleteCmd.Flags().UintVarP(&line, "line", "l", 0, "Line number to delete")
	RootCmd.AddCommand(deleteCmd)
}

// 接受一个参数m 删除行号为m的数据
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除指定行数据",
	Long:  "删除指定行数据 可以先使用 list 查看行号",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if line <= 0 {
			os.Exit(0)
		}
		err := funcs.Delete(line)
		if err != nil {
			fmt.Printf("删除错误 err: %v", errors.Unwrap(err))
		}
	},
}
