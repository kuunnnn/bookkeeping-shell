package del

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var line uint

func Register(root *cobra.Command) {
	command := createCommand()
	command.Flags().UintVarP(&line, "line", "l", 0, "Line number to delete")
	root.AddCommand(command)
}
func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "删除指定行数据",
		Long:  "删除指定行数据 可以先使用 list 查看行号",
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if line <= 0 {
				os.Exit(0)
			}
			err := delete(line)
			if err != nil {
				fmt.Printf("删除错误 err: %v", errors.Unwrap(err))
			}
		},
	}
}
