package main

import (
	"fmt"
	"os"

	"github.com/hiracy/zabton/logger"
	"github.com/hiracy/zabton/zabbix"
	"github.com/urfave/cli"
)

const (
	ENV_ZABBIX_API_URL      = "ZABTON_ZABBIX_URL"
	ENV_ZABBIX_API_USER     = "ZABTON_ZABBIX_USER"
	ENV_ZABBIX_API_PASSWORD = "ZABTON_ZABBIX_PASSWORD"
)

// Commands cli.Command object list
var Commands = []cli.Command{
	apiInfoCmd,
	pullCmd,
	pushCmd,
	diffCmd,
}

var apiInfoCmd = cli.Command{
	Name:  "info",
	Usage: "show zabbix server api version",
	Description: `
                Show Zabbix Server API Version.
`,
	Action: doApiInfoCmd,
}

var pullCmd = cli.Command{
	Name:  "pull",
	Usage: "pull configs from zabbix server",
	Description: `
                Pull text config file from specified Zabbix Server.
`,
	Action: doPullCmd,
}

var pushCmd = cli.Command{
	Name:  "push",
	Usage: "push configs to zabbix server",
	Description: `
                Push text config file to specified Zabbix Server.
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

func doApiInfoCmd(c *cli.Context) error {
	logger.SetLevel(c.GlobalString("log-level"))

	var server string

	if server = os.Getenv(ENV_ZABBIX_API_URL); server == "" {
		server = c.GlobalString("server")
	}

	logger.Log("info", "start info cmd: "+
		"server="+server)

	if server == "" {
		logger.Log("warn", "--server(-s) arg is required.")
		return nil
	}

	api := zabbix.NewAPI(server, "", "")
	version, err := api.Version()

	if err != nil {
		logger.Log("error", "ApiInfo: "+err.Error())
		return nil
	}

	fmt.Println(version)
	return nil
}

func doPullCmd(c *cli.Context) error {
	logger.SetLevel(c.GlobalString("log-level"))

	token := login(c, "pull")
	logger.Log("debug", "auth: "+token)
	return nil
}

func doPushCmd(c *cli.Context) error {
	logger.SetLevel(c.GlobalString("log-level"))

	token := login(c, "push")
	logger.Log("debug", "auth: "+token)
	return nil
}

func doDiffCmd(c *cli.Context) error {
	logger.SetLevel(c.GlobalString("log-level"))

	token := login(c, "diff")
	logger.Log("debug", "auth: "+token)
	return nil
}

func login(c *cli.Context, mode string) string {
	var server string
	var user string
	var password string

	if server = os.Getenv(ENV_ZABBIX_API_URL); server == "" {
		server = c.GlobalString("server")
	}
	if user = os.Getenv(ENV_ZABBIX_API_USER); user == "" {
		user = c.GlobalString("user")
	}
	if password = os.Getenv(ENV_ZABBIX_API_PASSWORD); password == "" {
		password = c.GlobalString("password")
	}

	logger.Log("info", "start "+mode+" cmd: "+
		"server="+server)

	if server == "" || user == "" || password == "" {
		logger.Log("warn", "--server(-s) and --user(-u) and --password(-p) args are required.")
		return ""
	}

	api := zabbix.NewAPI(server, user, password)

	auth, err := api.Login()

	if err != nil {
		logger.Log("error", "Login: "+err.Error())
		return ""
	}

	logger.Log("debug", "auth: "+auth)
	return auth
}
