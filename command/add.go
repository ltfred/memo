package command

import (
	"errors"
	"time"

	"github.com/ltfred/memo/constant"

	"github.com/ltfred/memo/pkg/parser"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

var Add cli.Command

func init() {
	Add = cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add a record",
		Action:  addAction,
	}
}

func addAction(c *cli.Context) error {
	var (
		name     string
		date     string
		content  string
		priority string
		err      error
	)

	prompt := promptui.Prompt{Label: "Name", Validate: func(s string) error {
		if s == "" {
			return errors.New("name cannot be empty")
		}

		return nil
	}}
	if name, err = prompt.Run(); err != nil {
		return err
	}

	prompt = promptui.Prompt{Label: "Date", Validate: func(s string) error {
		if s == "" {
			return errors.New("date cannot be empty")
		}

		return nil
	}}
	if date, err = prompt.Run(); err != nil {
		return err
	}

	prompt = promptui.Prompt{Label: "Content"}
	if content, err = prompt.Run(); err != nil {
		return err
	}

	sel := promptui.Select{Label: "Priority", Items: constant.MemoPriorityValues}
	if _, priority, err = sel.Run(); err != nil {
		return err
	}

	par := parser.Parser{}

	return par.Add(parser.Memo{
		Name:     name,
		Date:     date,
		Content:  content,
		Priority: priority,
		Status:   constant.MemoStatusUndo,
		CreateAt: time.Now().Unix(),
	})
}
