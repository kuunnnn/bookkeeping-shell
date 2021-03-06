package list

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var limit uint
var offset uint

func Register(root *cobra.Command)  {
	command := createCommand()
	command.Flags().UintVarP(&limit, "limit", "l", 10, "Number of data to be displayed")
	command.Flags().UintVarP(&offset, "offset", "o", 0, "Offset")
	root.AddCommand(command)
}

func handler(cmd *cobra.Command, args []string) {
	fmt.Println(limit)
	if limit <= 0 {
		limit = 10
	}
	err := Cat(limit, offset)
	if err != nil {
		fmt.Printf("打印数据失败 err: %v\n", errors.Unwrap(err))
	}
}


// 接受一个参数m 输出 最近m 条数据带行号
func createCommand()*cobra.Command  {
	return &cobra.Command{
		Use:   "list",
		Short: "查看最近[]行数据",
		Long: `查看最近[]行数据 默认 10
注意 
	1. 从文件末行开始找到第一个不是空行的开始计数
	2. 中间的空行会被跳过
	3. 空行不会被打印
**一行的内容只有\n 表示空行**
`,
		Args: cobra.MaximumNArgs(0),
		Run: handler,
	}
}


