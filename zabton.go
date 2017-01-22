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

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)

	app.Run(os.Args)
}
