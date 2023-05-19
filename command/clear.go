package command

import (
	"github.com/ltfred/memo/pkg/parser"
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
	par := parser.Parser{}
	return par.Clear()
}
