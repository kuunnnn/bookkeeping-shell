package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(viewCmd)
}

// 输出指定格式的文件
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "展示图表",
	Long:  "在浏览器中展示图表",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
		//if line == 0 {
		//	os.Exit(0)
		//}
		//err := _func.Delete(line)
		//if err != nil {
		//	fmt.Printf("删除错误 err: %v", errors.Cause(err))
		//	os.Exit(-1)
		//}
	},
}
