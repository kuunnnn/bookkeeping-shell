package backup

import (
	"bookkeeping-shell/funcs"
	read_data "bookkeeping-shell/read-data"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(createCommand())
}

func handler(cmd *cobra.Command, args []string) {
	err := read_data.Backup()
	if err != nil {
		fmt.Printf("备份数据失败 err: %v", errors.Unwrap(err))
	}
}
func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "backup",
		Short: "备份数据",
		Long:  fmt.Sprintf(`备份数据到[ %s ]目录下,且会添加日期到文件名后`, funcs.FileBackupPath),
		Args:  cobra.MaximumNArgs(0),
		Run:   handler,
	}
}
