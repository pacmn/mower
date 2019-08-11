package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "mower"
	app.Usage = "mower: app to find out where your mower ended up"
	app.Version = "1.0.0"

	registerCommands(app)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
