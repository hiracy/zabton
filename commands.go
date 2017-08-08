package main

import (
	"github.com/hiracy/zabton/logger"
	"github.com/hiracy/zabton/zabbix"
	"github.com/urfave/cli"
)

// Commands cli.Command object list
var Commands = []cli.Command{
	pullCmd,
	pushCmd,
	diffCmd,
}

var pullCmd = cli.Command{
	Name:  "pull",
	Usage: "pull configs from zabbix server",
	Description: `
                Pull text config file from specified Zabbix server.
`,
	Action: doPullCmd,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "server, s",
			Usage: "Zabbix server url(ex: http://api.zabbix.zabton.jp/api_jsonrpc.php)"},
		cli.StringFlag{
			Name:  "user, u",
			Usage: "Login user"},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Login password"},
	},
}

var pushCmd = cli.Command{
	Name:  "push",
	Usage: "push configs to zabbix server",
	Description: `
                Push text config file to specified Zabbix server.
`,
	Action: doPushCmd,
}

var diffCmd = cli.Command{
	Name:  "diff",
	Usage: "show configs difference",
	Description: `
                Show configs difference between zabbix server and local file
`,
	Action: doPushCmd,
}

func doPullCmd(c *cli.Context) error {
	logger.SetLevel(c.GlobalString("log-level"))

	server := c.String("server")
	user := c.String("user")
	password := c.String("password")

	logger.Log("info", "start pull cmd: "+
		"server="+server)

	if server == "" || user == "" || password == "" {
		logger.Log("warn", "--server(-s) and --user(-u) and --password(-p) args are required.")
		return nil
	}

	api := zabbix.NewAPI(server, user, password)

	auth, err := api.Login()

	if err != nil {
		logger.Log("error", "Login: "+err.Error())
		return nil
	}

	logger.Log("debug", "auth: "+auth)

	return nil
}

func doPushCmd(c *cli.Context) error {
	logger.Log("info", "start push cmd")
	return nil
}

func doDiffCmd(c *cli.Context) error {
	logger.Log("info", "start diff cmd")
	return nil
}
