package cmd

import (
	"bookkeeping-shell/funcs"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var target string
var fromCategoryName []string
var tagList bool

func init() {
	tagCmd.Flags().BoolVarP(&tagList, "list", "l", false, "列出类别列表")
	tagCmd.Flags().StringSliceVarP(&fromCategoryName, "modify", "m", nil, "修改前的名称")
	tagCmd.Flags().StringVarP(&target, "target", "t", "", "修改后的名称")
	RootCmd.AddCommand(tagCmd)
}

// 输出指定格式的文件
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "对类别进行操作",
	Long: `
--merge  合并 2 个类别数据 
--list   列出类别列表
--modify 修改类别名称

`,
	Args: cobra.MaximumNArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if tagList && target == "" && fromCategoryName == nil {
			err = funcs.PrintCategoryList()
		} else if fromCategoryName != nil && target != "" {
			err = funcs.ModifyCategoryName(fromCategoryName, target)
			err = funcs.PrintCategoryList()
		}
		if err != nil {
			fmt.Printf("%v\n", errors.Unwrap(err))
		}
	},
}
