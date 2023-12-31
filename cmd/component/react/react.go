/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package react

import (
	"fmt"

	"github.com/hoehwa/but/cmd/component"
	"github.com/spf13/cobra"
)

// reactCmd represents the react command.
var reactCmd = &cobra.Command{
	Use:   "react",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		Run 'but component react <subcommand>' for more information on a specific command.
		
		you can use following subcommands:
		- but component react display
		- but component react feedback
		- but component react hooks
		- but component react input
		- but component react stateful
		- but component react visuals
		`)
	},
}

func init() {
	component.ComponentCmd.AddCommand(reactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reactCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
