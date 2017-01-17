package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

var hostCmd = []cli.Command{
	{
		Name:  "host",
		Usage: "operate host api",
		Subcommands: []cli.Command{
			{
				Name:  "get",
				Usage: "The method allows to retrieve hosts according to the given parameters.",
				Action: func(c *cli.Context) error {
					fmt.Println("new task template: ", c.Args().First())
					return nil
				},
			},
		},
	},
}
