/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package patterns

import (
	"fmt"

	"github.com/hoehwa/but/cmd/fiddle"
	"github.com/spf13/cobra"
)

// patternsCmd represents the patterns command
var patternsCmd = &cobra.Command{
	Use:   "patterns",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		Run 'but fiddle pattern <subcommand>' for more information on a specific command.
		
		you can use following subcommands:
		- but fiddle patterns animation
		- but fiddle patterns clipboard
		- but fiddle patterns components
		- but fiddle patterns exampleSet
		- but fiddle patterns files
		- but fiddle patterns layout
		- but fiddle patterns media
		- but fiddle patterns theming
		- but fiddle patterns webApps
		- but fiddle patterns webVitalPtrns
		`)
	},
}

func init() {
	fiddle.FiddleCmd.AddCommand(patternsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patternsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patternsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
