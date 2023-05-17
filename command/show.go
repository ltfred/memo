package command

import (
	"github.com/ltfred/memo/pkg/parser"
	"github.com/ltfred/memo/types"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"os"
	"sort"
)

var Show cli.Command

func init() {
	Add = cli.Command{
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "show all record",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "s, sort",
				Usage: "sort record with value, value: [p,priority; n,name; d,date]",
			},
		},
		Action: showAction,
	}
}

func showAction(c *cli.Context) error {
	par := parser.Parser{}
	memos, err := par.Show()
	if err != nil {
		return err
	}

	sort.Slice(memos, func(i, j int) bool {
		return memos[i].Priority < memos[j].Priority
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Date", "Content", "Priority"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	table.SetRowLine(true)

	for i := range memos {
		colors := []tablewriter.Colors{{}, {}, {}, {}}
		if memos[i].Priority == types.MEMO_PRIORITY__IMPORTANT {
			colors[3] = tablewriter.Colors{tablewriter.FgRedColor}
		}
		table.Rich([]string{memos[i].Name, memos[i].Date, memos[i].Content, memos[i].Priority.String()}, colors)
	}

	table.Render()

	return nil
}
