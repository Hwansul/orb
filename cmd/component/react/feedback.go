/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package react

import (
	"github.com/hoehwa/but/internal"
	"github.com/spf13/cobra"
)

// feedbackCmd represents the feedback command.
var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "Open stackblitz sandbox for react component to send feedbacks to users.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitComponents("/component/react/feedback/")
	},
}

func init() {
	reactCmd.AddCommand(feedbackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// feedbackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// feedbackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
