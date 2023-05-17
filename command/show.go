package command

import (
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ltfred/memo/pkg/parser"
	"github.com/ltfred/memo/types"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

var Show cli.Command

func init() {
	Show = cli.Command{
		Name:    "show",
		Aliases: []string{"s"},
		Usage:   "show all record",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "s, sort",
				Usage: "sort record with value, value: [p(priority),[asc,desc]; d(date),[asc,desc]]",
			},
		},
		Action: showAction,
	}
}

func showAction(c *cli.Context) error {
	par := parser.Parser{}
	memosMap, err := par.Show()
	if err != nil {
		return err
	}
	memos := make([]parser.Memo, 0, len(memosMap))
	for k := range memosMap {
		memos = append(memos, memosMap[k])
	}

	sort.Slice(memos, func(i, j int) bool {
		s := strings.Split(c.String("sort"), ",")
		switch len(s) {
		case 0:
		case 1:
			switch s[0] {
			case "p", "priority":
				return memos[i].Priority < memos[j].Priority
			case "d", "date":
				return memos[i].Date > memos[j].Date
			}
		case 2:
			switch s[0] {
			case "p", "priority":
				if s[1] == "asc" {
					return memos[i].Priority > memos[j].Priority
				}
				return memos[i].Priority < memos[j].Priority
			case "d", "date":
				if s[1] == "asc" {
					return memos[i].Date < memos[j].Date
				}
				return memos[i].Date > memos[j].Date
			}
		}

		return memos[i].Date > memos[j].Date
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Num", "Uuid", "Name", "Date", "Content", "Priority", "Status", "CreateAt"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	table.SetRowLine(true)
	table.SetCaption(true, strings.Join([]string{strconv.Itoa(len(memos)), "Total."}, " "))
	for i := range memos {
		colors := []tablewriter.Colors{{}, {}, {}, {}, {}, {}, {}}
		if memos[i].Priority == types.MemoPriorityHigh {
			colors[5] = tablewriter.Colors{tablewriter.FgRedColor}
		}
		table.Rich([]string{strconv.Itoa(i + 1), strconv.FormatInt(memos[i].CreateAt, 10), memos[i].Name,
			memos[i].Date, memos[i].Content, memos[i].Priority.String(), memos[i].Status.String(),
			time.Unix(memos[i].CreateAt, 0).Format("2006-01-02 15:04:05")}, colors)
	}

	table.Render()

	return nil
}
