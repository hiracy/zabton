package main

import (
	"fmt"
	"os"

	"github.com/hiracy/zabton/commands"
	"github.com/urfave/cli"
)

var Version string

func main() {
	cli.VersionPrinter = printVersion
	app := cli.NewApp()
	app.Name = "zabton"
	app.Usage = "Zabbix API and CLI tool set."
	app.Version = Version
	commands.Build(app)

	app.Run(os.Args)
}

func printVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%v\n", c.App.Version)
}
