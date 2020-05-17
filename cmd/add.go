package cmd

import (
	"bookkeeping-shell/funcs"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	//"strconv"
)

var typ string
var desc string
var day int
var money float64

func init() {
	addCmd.Flags().StringVarP(&typ, "type", "t", "e", "Income(i) or Expense(e)")
	addCmd.Flags().StringVarP(&desc, "desc", "d", "", "a description")
	addCmd.Flags().IntVarP(&day, "time", "o", 0, "0 is today -1 is yesterday")
	addCmd.Flags().Float64VarP(&money, "money", "m", 0, "money (required)")
	_ = addCmd.MarkFlagRequired("money")
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [category]",
	Short: "新增一条记录",
	Long:  `新增一条记录 (1 个参数: 类别) flag money 是必须的`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// cobra 会吧负数解析为 flag 虽然 money 不能被设置为负数,但还是换成了 flag 实现
		//money, err := strconv.ParseFloat(args[0], 10)
		//if err != nil {
		//	fmt.Printf("money 需要是一个可以被解析成 float64 的参数 [ %s ]", args[0])
		//	os.Exit(-1)
		//}
		tp := funcs.INCOME
		if typ == "e" {
			tp = funcs.EXPENSE
		}
		if desc == "" {
			desc = "无"
		}
		if money <= 0 {
			fmt.Printf("money[ %.2f ]不能小于等于0\n", money)
			os.Exit(0)
		}
		err := funcs.Add(money, args[0], desc, day, tp)
		if err != nil {
			fmt.Printf("插入数据失败 err: %v", errors.Unwrap(err))
		}
	},
}
