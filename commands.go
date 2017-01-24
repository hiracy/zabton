package main

import (
	"github.com/hiracy/zabton/logger"
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
	logger.Log("info", "start pull cmd")
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
