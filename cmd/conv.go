package cmd

import (
	"bookkeeping-shell/funcs"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var output string

func init() {
	convCmd.Flags().StringVarP(&output, "output", "o", "json", "Output to the specified format file")
	RootCmd.AddCommand(convCmd)
}

// 输出指定格式的文件
var convCmd = &cobra.Command{
	Use:   "conv",
	Short: "输出指定格式的文件",
	Long:  "输出指定格式的文件 暂时仅支持 [json]",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if output == "json" {
			err := funcs.ToJSON()
			if err != nil {
				fmt.Printf("err: %v", errors.Unwrap(err))
			}
		}else {
			fmt.Println("TODO")
		}
	},
}
