// Orb shortcut show tables for vscode and obsidian
// Copyright © 2023 Mindulle <mindullestudio@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/jipilmuk/orb/utills/table"

	"github.com/spf13/cobra"
)

var dataArr = [][][]string{
	table.VscodeBody,
	table.ObsidianBody,
}

// shortcutCmd represents the shortcut command.
var shortcutCmd = &cobra.Command{
	Use:   "shortcut",
	Short: "Show shortcuts that you should be familiar enough with your editor.",
	Long: `This commands show markdown tables which represents 
what features each commands have and short descriptions.
This command has flags for text editors(or IDE) which you prefer.
`,

	Run: func(cmd *cobra.Command, args []string) {

		if vscodeOption, _ := cmd.Flags().GetBool("vscode"); vscodeOption {
			table.Print(dataArr[0], table.Header, table.Footer)
		}
		if obsidianOption, _ := cmd.Flags().GetBool("obsidian"); obsidianOption {
			table.Print(dataArr[1], table.Header, table.Footer)
		}

		fmt.Println("You can see a shortcut tables for obsidian or vscode. please run this command again with --vscode or --obsidian")
	},
}

func init() {
	RootCmd.AddCommand(shortcutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shortcutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	shortcutCmd.Flags().BoolP("vscode", "", false, "Show shortcuts for vscode")
	shortcutCmd.Flags().BoolP("obsidian", "", false, "Show my custom shortcuts for obsidian")
}
