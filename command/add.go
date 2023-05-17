package command

import (
	"github.com/ltfred/memo/pkg/parser"
	"github.com/ltfred/memo/types"
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
	prompt := promptui.Prompt{
		Label: "Name",
	}
	name, err := prompt.Run()
	if err != nil {
		return err
	}

	prompt = promptui.Prompt{
		Label: "Date",
	}
	date, err := prompt.Run()
	if err != nil {
		return err
	}

	prompt = promptui.Prompt{
		Label: "Content",
	}
	content, err := prompt.Run()
	if err != nil {
		return err
	}

	sel := promptui.Select{
		Label: "Priority",
		Items: types.MemoPriorityValues,
	}
	_, priority, err := sel.Run()
	if err != nil {
		return err
	}

	par := parser.Parser{}
	err = par.Add(parser.Memo{
		Name:     name,
		Date:     date,
		Content:  content,
		Priority: types.ParseMemoPriorityFromString(priority),
	})

	return err
}
