package view

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command)  {
	root.AddCommand(createCommand())
}

func createCommand()*cobra.Command  {
	return &cobra.Command{
		Use:   "view",
		Short: "展示图表",
		Long:  "在浏览器中展示图表",
		Args:  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			err := genViewJS()
			if err != nil {
				fmt.Printf("error: %v:\n", errors.Unwrap(err))
			}
			fmt.Println("正在打开浏览器 ...")
		},
	}
}
