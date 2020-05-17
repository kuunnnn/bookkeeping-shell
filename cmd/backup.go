package cmd

import (
	"bookkeeping-shell/funcs"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(backupCmd)
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "备份数据",
	Long:  fmt.Sprintf(`备份数据到[ %s ]目录下,且会添加日期到文件名后`,funcs.FileBackupPath),
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := funcs.Backup()
		if err != nil {
			fmt.Printf("备份数据失败 err: %v", errors.Unwrap(err))
		}
	},
}
