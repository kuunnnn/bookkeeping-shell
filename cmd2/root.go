package cmd2

import (
	"bookkeeping-shell/cmd2/add"
	"bookkeeping-shell/cmd2/analyze"
	"bookkeeping-shell/cmd2/backup"
	"bookkeeping-shell/cmd2/conv"
	"bookkeeping-shell/cmd2/del"
	"bookkeeping-shell/cmd2/list"
	"bookkeeping-shell/cmd2/sort"
	"bookkeeping-shell/cmd2/tag"
	"bookkeeping-shell/cmd2/view"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "bookkeeping",
	Short: "Hu Kun 的记账小程序",
	Long:  `Hu Kun 的记账小程序 存储为一个 .txt 数据文件`,
}

func init() {
	tag.Register(RootCmd)
	sort.Register(RootCmd)
	backup.Register(RootCmd)
	conv.Register(RootCmd)
	list.Register(RootCmd)
	analyze.Register(RootCmd)
	view.Register(RootCmd)
	add.Register(RootCmd)
	del.Register(RootCmd)
}
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
