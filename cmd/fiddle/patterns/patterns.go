/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package patterns

import (
	"fmt"

	"github.com/hoehwa/but/cmd/fiddle"
	"github.com/spf13/cobra"
)

// patternCmd represents the pattern command.
var patternCmd = &cobra.Command{
	Use:   "pattern",
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
		- but fiddle pattern animation
		- but fiddle pattern clipboard
		- but fiddle pattern components
		- but fiddle pattern exampleSet
		- but fiddle pattern files
		- but fiddle pattern layout
		- but fiddle pattern media
		- but fiddle pattern theming
		- but fiddle pattern webApps
		- but fiddle pattern webVitalPtrns
		`)
	},
}

func init() {
	fiddle.FiddleCmd.AddCommand(patternCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patternCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patternCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
