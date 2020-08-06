package sort

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command) {
	root.AddCommand(register())
}

func register() *cobra.Command {
	return &cobra.Command{
		Use:   "sort",
		Short: "排序",
		Long:  "对 text 文件数据进行排序",
		Args:  cobra.MaximumNArgs(0),
		Run:   handler,
	}
}

func handler(cmd *cobra.Command, args []string) {
	err := Sort();
	if err != nil {
		fmt.Printf("error: %v:\n", errors.Unwrap(err))
	}
}
