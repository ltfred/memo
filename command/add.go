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
	var (
		name     string
		date     string
		content  string
		priority string
		err      error
	)

	prompt := promptui.Prompt{Label: "Name"}
	if name, err = prompt.Run(); err != nil {
		return err
	}
	prompt = promptui.Prompt{Label: "Date"}
	if date, err = prompt.Run(); err != nil {
		return err
	}
	prompt = promptui.Prompt{Label: "Content"}
	if content, err = prompt.Run(); err != nil {
		return err
	}

	sel := promptui.Select{Label: "Priority", Items: types.MemoPriorityValues}
	if _, priority, err = sel.Run(); err != nil {
		return err
	}
	par := parser.Parser{}

	return par.Add(parser.Memo{
		Name:     name,
		Date:     date,
		Content:  content,
		Priority: types.ParseMemoPriorityFromString(priority),
	})
}
