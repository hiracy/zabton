package main

import (
	"fmt"
	"os"

	"github.com/hiracy/zabton/commands"
	"github.com/urfave/cli"
)

// Version is the package version
var Version string

func main() {
	cli.VersionPrinter = printVersion
	app := cli.NewApp()
	app.Name = "zabton"
	app.Usage = `CLI for managing Zabbix with text base config.

See Zabbix Official Documents(https://www.zabbix.com/documentation/3.0/manual/api)
	`
	app.Version = Version
	commands.Build(app)

	app.Run(os.Args)
}

func printVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%v\n", c.App.Version)
}
