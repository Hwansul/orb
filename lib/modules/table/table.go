package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Print prints a table with the given data, header, and footer.
//
// data: the data to be displayed in the table, assumes 2D slice of strings.
// header: the header row of the table, assumes slice of strings.
// footer: the footer row of the table, assumes slice of strings.
func Print(data [][]string, header []string, footer []string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	table.SetFooter(footer)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiWhiteColor, tablewriter.Bold, tablewriter.BgGreenColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor})

	table.SetFooterColor(
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.FgHiRedColor})

	table.AppendBulk(data)
	table.Render()
}
