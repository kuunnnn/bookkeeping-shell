package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(tagCmd)
}

// 输出指定格式的文件
var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "对类别进行操作",
	Long: `
--merge 合并 2 个标签数据 
`,
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}
