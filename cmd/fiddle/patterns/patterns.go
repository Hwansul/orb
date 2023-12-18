/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package patterns

import (
	"fmt"

	"github.com/jipilmuk/orb/cmd/fiddle"
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
		Run 'orb fiddle pattern <subcommand>' for more information on a specific command.
		
		you can use following subcommands:
		- orb fiddle pattern animation
		- orb fiddle pattern clipboard
		- orb fiddle pattern components
		- orb fiddle pattern exampleSet
		- orb fiddle pattern files
		- orb fiddle pattern layout
		- orb fiddle pattern media
		- orb fiddle pattern theming
		- orb fiddle pattern webApps
		- orb fiddle pattern webVitalPtrns
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
