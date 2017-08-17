package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/hiracy/zabton/logger"
	"github.com/hiracy/zabton/zabbix"
	"github.com/urfave/cli"
)

const (
	ENV_ZABBIX_API_URL      = "ZABTON_ZABBIX_URL"
	ENV_ZABBIX_API_USER     = "ZABTON_ZABBIX_USER"
	ENV_ZABBIX_API_PASSWORD = "ZABTON_ZABBIX_PASSWORD"
)

// AvailableObjects are list of available objects
var AvailableObjects = []string{
	"host",
	"hostgroup",
}

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
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "objects, o",
			Usage: "Apply specified objects only(comma or space separated)"},
	},
}

var pushCmd = cli.Command{
	Name:  "push",
	Usage: "push configs to zabbix server",
	Description: `
                Push text config file to specified Zabbix Server.
`,
	Action: doPushCmd,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "objects, o",
			Usage: "Apply specified objects only"},
	},
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

	objects := parseSubCmdArgs(c)

	api := login(c, "pull")

	if api == nil {
		return nil
	}

	client := NewClient(objects, api)

	for _, obj := range objects {
		ret := reflect.ValueOf(client).MethodByName("Pull" + strings.Title(obj)).Call(nil)

		var err error
		if ret[0].Interface() != nil {
			err = ret[0].Interface().(error)
		}

		if err != nil {
			logger.Log("error", "pull"+strings.Title(obj)+": "+err.Error())
			return nil
		}
	}

	return nil
}

func doPushCmd(c *cli.Context) error {
	logger.SetLevel(c.GlobalString("log-level"))

	api := login(c, "push")

	if api == nil {
		return nil
	}

	logger.Log("debug", "auth: "+api.Auth)
	return nil
}

func doDiffCmd(c *cli.Context) error {
	logger.SetLevel(c.GlobalString("log-level"))

	api := login(c, "diff")

	if api == nil {
		return nil
	}

	logger.Log("debug", "auth: "+api.Auth)
	return nil
}

func parseSubCmdArgs(c *cli.Context) (objects []string) {
	argObj := c.String("objects")
	var parsedObj []string

	if argObj == "" {
		parsedObj = AvailableObjects
	} else if strings.Index(argObj, " ") > 0 {
		parsedObj = strings.Split(argObj, " ")
	} else if strings.Index(argObj, ",") > 0 {
		parsedObj = strings.Split(argObj, ",")
	} else {
		parsedObj = []string{argObj}
	}

	for _, parsed := range parsedObj {
		for _, avails := range AvailableObjects {
			if avails == parsed {
				objects = append(objects, parsed)
			}
		}
	}

	return objects
}

func login(c *cli.Context, mode string) *zabbix.API {
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
		return nil
	}

	api := zabbix.NewAPI(server, user, password)

	auth, err := api.Login()

	logger.Log("debug", "auth: "+auth)

	if err != nil {
		logger.Log("error", "Login: "+err.Error())
		return nil
	}

	return api
}
