package commands

import (
	"github.com/urfave/cli"
)

// Build creates the command objects
func Build(app *cli.App) {
	commands := []cli.Command{}
	commands = append(commands, hostCmd...)
	app.Commands = commands
}
