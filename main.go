package main

import (
	"log"
	"os"

	"github.com/ltfred/memo/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.App{
		Name:  "memo",
		Usage: "This is a command line memo.",
	}
	app.Commands = cli.Commands{command.Add, command.Show, command.Delete, command.Modify}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
