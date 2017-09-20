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
			Name:  "log-level, l",
			Value: "warn",
			Usage: "output log level[trace|debug|info|warn|fatal|alert]",
		},
		cli.StringFlag{
			Name:  "server, s",
			Usage: "Zabbix server url(ex: http://api.zabbix.zabton.jp/api_jsonrpc.php)"},
		cli.StringFlag{
			Name:  "user, u",
			Usage: "Login user"},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Login password"},
		cli.StringFlag{
			Name:  "config, c",
			Value: "zabton.conf",
			Usage: "Zabton definitions file path"},
	}

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)

	app.Run(os.Args)
}
