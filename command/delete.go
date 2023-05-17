package command

import (
	"github.com/ltfred/memo/pkg/parser"
	"github.com/urfave/cli"
)

var Delete cli.Command

func init() {
	Delete = cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete record",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "u",
				Usage:    "delete record by uuid",
				Required: true,
			},
		},
		Action: deleteAction,
	}
}

func deleteAction(c *cli.Context) error {
	par := parser.Parser{}

	return par.Delete(c.String("u"))
}
