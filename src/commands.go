package main

import (
	"github.com/urfave/cli"

	"mower/mow"
)

func registerCommands(app *cli.App) {
	app.Commands = []cli.Command{
		mow.Command,
	}
}
