package command

import (
	"github.com/ltfred/memo/pkg/parser"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

var Clear cli.Command

func init() {
	Clear = cli.Command{
		Name:    "clear",
		Aliases: []string{"c"},
		Usage:   "clear all records",
		Action:  clearAction,
	}
}

func clearAction(c *cli.Context) error {
	prompt := promptui.Prompt{Label: "clear all records", IsConfirm: true}
	_, err := prompt.Run()
	if err != nil {
		return err
	}

	par := parser.Parser{}
	return par.Clear()
}
