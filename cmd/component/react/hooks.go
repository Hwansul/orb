/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package react

import (
	"github.com/hoehwa/but/utills"
	"github.com/spf13/cobra"
)

// hooksCmd represents the hooks command.
var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Open stackblitz sandbox for react component with custom hook to delegate complicated tasks to react component.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.InitComponents("/component/react/hooks/")
	},
}

func init() {
	reactCmd.AddCommand(hooksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hooksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hooksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
