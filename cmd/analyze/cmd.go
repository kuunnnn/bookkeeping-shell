package analyze

import (
	read_data "bookkeeping-shell/read-data"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var species string
var billType string

func Register(root *cobra.Command) {
	command := createCommand()
	command.Flags().
		StringVarP(&species, "species", "s", "e", "Income(i) or Expense(e)")
	command.Flags().
		StringVarP(&billType, "type", "t", "year", "year or month or tag")
	root.AddCommand(command)
}
func handler(cmd *cobra.Command, args []string) {
	s := read_data.INCOME
	if species != "i" {
		s = read_data.EXPENSE
	}
	t := MONTH
	if billType != "month" && billType != "tag" {
		t = YEAR
	} else if billType == "month" {
		t = MONTH
	} else {
		t = TAG
	}
	err := OutputBillData(t, s)
	if err != nil {
		fmt.Printf("%v\n", errors.Unwrap(err))
	}
}
func createCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "analyze",
		Short: "data analysis",
		Long:  "temporary data by year month and tag",
		Args:  cobra.MaximumNArgs(0),
		Run:   handler,
	}
}
