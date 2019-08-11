package mow

import (
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:    "find",
	Usage:   "Find where did the mower ended up",
	Aliases: []string{"fi"},
	Action:  findMower,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "file, f",
			Usage:    "Find where the mower ended up with data from file `FILE`",
			Required: true,
		},
	},
}
