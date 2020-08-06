package cmd

import (
	"bookkeeping-shell/cmd/add"
	"bookkeeping-shell/cmd/analyze"
	"bookkeeping-shell/cmd/backup"
	"bookkeeping-shell/cmd/conv"
	"bookkeeping-shell/cmd/del"
	"bookkeeping-shell/cmd/list"
	"bookkeeping-shell/cmd/sort"
	"bookkeeping-shell/cmd/tag"
	"bookkeeping-shell/cmd/view"
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
