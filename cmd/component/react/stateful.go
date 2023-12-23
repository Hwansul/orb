/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package react

import (
	"github.com/hoehwa/but/utills"
	"github.com/spf13/cobra"
)

// statefulCmd represents the stateful command.
var statefulCmd = &cobra.Command{
	Use:   "stateful",
	Short: "Open stackblitz sandbox for react stateful component.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.InitComponents("/component/react/stateful/")
	},
}

func init() {
	reactCmd.AddCommand(statefulCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statefulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statefulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
