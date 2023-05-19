package command

import (
	"errors"
	"fmt"

	"github.com/ltfred/memo/constant"

	"github.com/ltfred/memo/pkg/parser"
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

	sel := promptui.Select{Label: fmt.Sprintf("Status(%s)", record.Status), Items: constant.MemoStatusValues}
	if _, record.Status, err = sel.Run(); err != nil {
		return err
	}

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

	prompt := promptui.Prompt{Label: fmt.Sprintf("Name(%s)", record.Name), Default: record.Name}
	if record.Name, err = prompt.Run(); err != nil {
		return err
	}

	prompt = promptui.Prompt{Label: fmt.Sprintf("Date(%s)", record.Date), Default: record.Date}
	if record.Date, err = prompt.Run(); err != nil {
		return err
	}

	prompt = promptui.Prompt{Label: "Content", Default: record.Content}
	if record.Content, err = prompt.Run(); err != nil {
		return err
	}

	sel := promptui.Select{Label: fmt.Sprintf("Priority(%s)", record.Priority), Items: constant.MemoPriorityValues}
	if _, record.Priority, err = sel.Run(); err != nil {
		return err
	}

	return par.Modify(uuid, record)
}
