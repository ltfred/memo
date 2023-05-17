package command

import "github.com/urfave/cli"

var Delete cli.Command

func init() {
	Delete = cli.Command{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete record",
		Action:  deleteAction,
	}
}

func deleteAction(c *cli.Context) error {
	return nil
}
