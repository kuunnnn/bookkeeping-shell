package tag

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)


var targetCategoryName string
var fromCategoryName []string

func Register(root *cobra.Command) {
	cmd := createCommand()
	cmd.Flags().StringSliceVarP(&fromCategoryName, "modify", "m", nil, "修改前的名称")
	cmd.Flags().StringVarP(&targetCategoryName, "target", "t", "", "修改后的名称")
	root.AddCommand(cmd)
}

func handler(cmd *cobra.Command, args []string) {
	var err error
	if fromCategoryName != nil && targetCategoryName != "" {
		err = ModifyCategoryName(fromCategoryName, targetCategoryName)
		err = PrintCategoryList()
	} else {
		err = PrintCategoryList()
	}
	if err != nil {
		fmt.Printf("%v\n", errors.Unwrap(err))
	}
}

func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tag",
		Short: "对类别进行操作",
		Long: `
--modify 要修改的 Category Name  []string   例: -m="A,B" -m "A"
--target 修改后的 Category Name  string
`,
		Args: cobra.MaximumNArgs(5),
		Run:  handler,
	}
}
