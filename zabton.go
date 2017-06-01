package main

import (
	"os"
	"runtime"

	"github.com/urfave/cli"
)

func main() {
	cli.VersionPrinter = printVersion
	app := cli.NewApp()
	app.Name = "zabton"
	app.Usage = "CLI tool for managing Zabbix with text base config."
	app.Version = Version
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "apifile",
			Value: "",
			Usage: "another api defined config file",
		},
		cli.StringFlag{
			Name:  "log-level",
			Value: "info",
			Usage: "output log level",
		},
	}

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)

	app.Run(os.Args)
}
