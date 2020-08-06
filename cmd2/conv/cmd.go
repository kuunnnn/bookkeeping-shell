package conv

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var output string

func Register(command *cobra.Command) {
	cmd := createCommand()
	cmd.Flags().StringVarP(&output, "output", "o", "json", "Output to the specified format file")
	command.AddCommand(cmd)
}
func handler(cmd *cobra.Command, args []string) {
	if output == "json" {
		err := ToJSON()
		if err != nil {
			fmt.Printf("err: %v", errors.Unwrap(err))
		}
	} else {
		fmt.Println("TODO")
	}
}
func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "conv",
		Short: "输出指定格式的文件",
		Long:  "输出指定格式的文件 暂时仅支持 [json]",
		Args:  cobra.MaximumNArgs(0),
		Run:   handler,
	}
}
