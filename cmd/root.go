package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var RootCmd = &cobra.Command{
	Use:   "bookkeeping",
	Short: "Hu Kun 的记账小程序",
	Long: `Hu Kun 的记账小程序 存储为一个 .txt 数据文件`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// 可以加载一个配置文件
//var cfgFile string
//funcs init() {
//	cobra.OnInitialize(initConfig)
//	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bookkeeping.yaml)")
//}
//
//funcs initConfig() {
//	if cfgFile != "" {
//		viper.SetConfigFile(cfgFile)
//	}
//	viper.SetConfigName(".bookkeeping")
//	viper.AddConfigPath("$HOME")
//	viper.AutomaticEnv()
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}
