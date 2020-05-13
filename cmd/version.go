package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const VERSION = "v1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: VERSION,
	Long:  VERSION,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
