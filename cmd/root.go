/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/footgod368/fy-cli/internal/cmd/common"
	"github.com/footgod368/fy-cli/internal/cmd/root"

	"github.com/spf13/cobra"
)

const longDescription = `fy 是一个命令行英文词典工具，支持查询单词释义、同义词辨析等功能。
示例：
- 查询单词：$ fy insist
- 查询单词及其同义词辨析：$ fy -d insist
- 查看帮助：$ fy -h
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fy [text]",
	Short: "命令行英文词典",
	Long:  longDescription,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return root.Complete(toComplete)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return root.PreRunE(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return root.RunE(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fy-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "Verbose mode")
	rootCmd.Flags().BoolVarP(&root.Discriminate, "discriminate", "d", false, "Discriminate words")
	rootCmd.Flags().StringVarP(&root.Query, "query", "q", "", "Query words")
}
