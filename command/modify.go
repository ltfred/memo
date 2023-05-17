package command

import (
	"errors"
	"fmt"
	"github.com/ltfred/memo/pkg/parser"
	"github.com/ltfred/memo/types"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

var Modify cli.Command

func init() {
	Modify = cli.Command{
		Name:    "modify",
		Aliases: []string{"m"},
		Usage:   "modify record",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "u, uuid",
				Usage: "modify record by uuid",
			},
		},
		Action: modifyAction,
		Subcommands: []cli.Command{
			{
				Name:    "status",
				Aliases: []string{"s"},
				Usage:   "modify record status by uuid",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "u, uuid",
						Usage:    "modify record status by uuid",
						Required: true,
					},
				},
				Action: modifyStatusAction,
			},
		},
	}
}

func modifyStatusAction(c *cli.Context) error {
	uuid := c.String("uuid")
	par := parser.Parser{}
	record, err := par.GetRecord(uuid)
	if err != nil {
		return err
	}

	sel := promptui.Select{Label: fmt.Sprintf("Status(%s)", record.Status.String()), Items: types.MemoStatusValues}
	_, status, err := sel.Run()
	if err != nil {
		return err
	}

	record.Status = types.ParseMemoStatusFromString(status)

	return par.Modify(uuid, record)
}

func modifyAction(c *cli.Context) error {
	uuid := c.String("u")
	if uuid == "" {
		return errors.New("should use `memo m -u <uuid>`")
	}

	par := parser.Parser{}
	record, err := par.GetRecord(uuid)
	if err != nil {
		return err
	}

	var (
		name     string
		date     string
		content  string
		priority string
	)

	prompt := promptui.Prompt{Label: fmt.Sprintf("Name(%s)", record.Name), Default: record.Name}
	if name, err = prompt.Run(); err != nil {
		return err
	}

	prompt = promptui.Prompt{Label: fmt.Sprintf("Date(%s)", record.Date), Default: record.Date}
	if date, err = prompt.Run(); err != nil {
		return err
	}

	prompt = promptui.Prompt{Label: "Content", Default: record.Content}
	if content, err = prompt.Run(); err != nil {
		return err
	}

	sel := promptui.Select{Label: fmt.Sprintf("Priority(%s)", record.Priority), Items: types.MemoPriorityValues}
	if _, priority, err = sel.Run(); err != nil {
		return err
	}

	record.Name = name
	record.Date = date
	record.Content = content
	record.Priority = types.ParseMemoPriorityFromString(priority)

	return par.Modify(uuid, record)
}
