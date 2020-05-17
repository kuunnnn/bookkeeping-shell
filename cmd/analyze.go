package cmd

import (
	"bookkeeping-shell/funcs"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var species string
var billType string

func init() {
	analyzeCmd.Flags().
		StringVarP(&species, "species", "s", "e", "Income(i) or Expense(e)")
	analyzeCmd.Flags().
		StringVarP(&billType, "type", "t", "year", "year or month or tag")
	RootCmd.AddCommand(analyzeCmd)
}

// 输出指定格式的文件
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "data analysis",
	Long:  "temporary data by year month and tag",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		s := funcs.INCOME
		if species != "i" {
			s = funcs.EXPENSE
		}
		t := funcs.MONTH
		if billType != "month" && billType != "tag" {
			t = funcs.YEAR
		} else if billType == "month" {
			t = funcs.MONTH
		} else {
			t = funcs.TAG
		}
		err := funcs.OutputBillData(t, s)
		if err != nil {
			fmt.Printf("%v\n", errors.Unwrap(err))
		}
	},
}
