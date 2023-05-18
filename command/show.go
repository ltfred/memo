package command

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ltfred/memo/constant"

	"github.com/ltfred/memo/pkg/parser"
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
				Name:  "p, priority",
				Usage: "filter record with priority",
			},
			cli.StringFlag{
				Name:  "s, status",
				Usage: "filter record with status",
			},
			cli.StringFlag{
				Name:  "o, order",
				Usage: "order by, value: [p(priority),[asc,desc]; d(date),[asc,desc]]",
			},
		},
		Action: showAction,
	}
}

func showAction(c *cli.Context) error {
	priority := strings.ToLower(c.String("priority"))
	if priority != "" {
		if _, ok := constant.MemoPriorityValuesMap[priority]; !ok {
			return errors.New(fmt.Sprintf("priority value should be in `%v`", constant.MemoPriorityValues))
		}
	}
	status := strings.ToLower(c.String("status"))
	if status != "" {
		if _, ok := constant.MemoStatusValuesMap[status]; !ok {
			return errors.New(fmt.Sprintf("status value should be in `%v`", constant.MemoStatusValues))
		}
	}

	par := parser.Parser{}
	memosMap, err := par.Show()
	if err != nil {
		return err
	}

	memos := make([]parser.Memo, 0, len(memosMap))
	for k := range memosMap {
		if priority != "" || status != "" {
			if priority != "" && status != "" {
				if memosMap[k].Priority == priority && memosMap[k].Status == status {
					memos = append(memos, memosMap[k])
				}
			} else if priority != "" {
				if memosMap[k].Priority == priority {
					memos = append(memos, memosMap[k])
				}
			} else if status != "" {
				if memosMap[k].Status == status {
					memos = append(memos, memosMap[k])
				}
			}
		} else {
			memos = append(memos, memosMap[k])
		}
	}

	sort.Slice(memos, func(i, j int) bool {
		s := strings.Split(c.String("order"), ",")
		switch len(s) {
		case 1:
			switch s[0] {
			case "p", "priority":
				return constant.MemoPriorityValuesMap[memos[i].Priority] < constant.MemoPriorityValuesMap[memos[j].Priority]
			case "d", "date":
				return memos[i].Date > memos[j].Date
			}
		case 2:
			switch s[0] {
			case "p", "priority":
				if s[1] == "asc" {
					return constant.MemoPriorityValuesMap[memos[i].Priority] > constant.
						MemoPriorityValuesMap[memos[j].Priority]
				}
				return constant.MemoPriorityValuesMap[memos[i].Priority] < constant.MemoPriorityValuesMap[memos[j].Priority]
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
		if memos[i].Priority == constant.MemoPriorityHigh {
			colors[5] = tablewriter.Colors{tablewriter.FgRedColor}
		}
		table.Rich([]string{strconv.Itoa(i + 1), strconv.FormatInt(memos[i].CreateAt, 10), memos[i].Name,
			memos[i].Date, memos[i].Content, memos[i].Priority, memos[i].Status,
			time.Unix(memos[i].CreateAt, 0).Format("2006-01-02 15:04:05")}, colors)
	}

	table.Render()

	return nil
}
