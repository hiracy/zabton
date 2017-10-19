package main

import (
	"fmt"
	"os"
	"os/user"
	"reflect"
	"strings"

	"github.com/hiracy/zabton/logger"
	"github.com/hiracy/zabton/zabbix"
	"github.com/urfave/cli"
)

const (
	envZabbixURL      = "ZABTON_ZABBIX_URL"
	envZabbixUser     = "ZABTON_ZABBIX_USER"
	envZabbixPassword = "ZABTON_ZABBIX_PASSWORD"
	envZabtonLogLevel = "ZABTON_LOG_LEVEL"
	envZabtonFilePath = "ZABTON_FILE_PATH"
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
	Action: doAPIInfoCmd,
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
		cli.StringFlag{
			Name:  "file, f",
			Usage: "File path of save destination."},
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
			Usage: "Apply specified objects only(comma or space separated)"},
		cli.StringFlag{
			Name:  "file, f",
			Usage: "File path of read destination."},
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

func doAPIInfoCmd(c *cli.Context) error {
	if logLevel := os.Getenv(envZabtonLogLevel); logLevel == "" {
		logger.SetLevel(c.GlobalString("log-level"))
	}

	var server string

	if server = os.Getenv(envZabbixURL); server == "" {
		server = c.GlobalString("server")
	}

	logger.Log("info", "start Version() for info cmd: "+
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

	objects, editables, writepath, err := parseCmdArgs(c)

	if writepath == "" {
		return nil
	}

	if err != nil {
		logger.Log("error", "parseCmdArgs: "+err.Error())
		return nil
	}

	api := login(c, "pull")

	if api == nil {
		return nil
	}

	client := NewClient(api, writepath, "", editables)

	for _, obj := range objects {
		ret := reflect.ValueOf(client).MethodByName("Pull" + strings.Title(obj)).Call(nil)

		if ret[0].Interface() != nil {
			err = ret[0].Interface().(error)
		}

		if err != nil {
			logger.Log("error", "Pull"+strings.Title(obj)+": "+err.Error())
			return nil
		}
	}

	return nil
}

func doPushCmd(c *cli.Context) error {

	objects, editables, readpath, err := parseCmdArgs(c)

	if readpath == "" {
		return nil
	}

	if err != nil {
		logger.Log("error", "parseCmdArgs: "+err.Error())
		return nil
	}

	api := login(c, "pull")

	if api == nil {
		return nil
	}

	client := NewClient(api, "", readpath, editables)

	for _, obj := range objects {
		ret := reflect.ValueOf(client).MethodByName("Push" + strings.Title(obj)).Call(nil)

		if ret[0].Interface() != nil {
			err = ret[0].Interface().(error)
		}

		if err != nil {
			logger.Log("error", "Push"+strings.Title(obj)+": "+err.Error())
			return nil
		}
	}

	return nil
}

func doDiffCmd(c *cli.Context) error {
	if logLevel := os.Getenv(envZabtonLogLevel); logLevel == "" {
		logger.SetLevel(c.GlobalString("log-level"))
	}

	api := login(c, "diff")

	if api == nil {
		return nil
	}

	logger.Log("debug", "auth: "+api.Auth)
	return nil
}

func parseCmdArgs(c *cli.Context) (objects []string, editables *EditableConfiguration, filepath string, err error) {

	if logLevel := os.Getenv(envZabtonLogLevel); logLevel == "" {
		logger.SetLevel(c.GlobalString("log-level"))
	}

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

	editables, err = LoadConfig(c.GlobalString("config"))

	if err != nil {
		return nil, nil, "", err
	}

	filepath = c.String("file")

	if filepath == "" {
		if filepath = os.Getenv(envZabtonFilePath); filepath == "" {
			logger.Log("warn", "--file(-f) arg is required.")
		}
	}

	usr, err := user.Current()

	if err != nil {
		return nil, nil, "", err
	}

	return objects, editables, strings.Replace(filepath, "~", usr.HomeDir, 1), nil
}

func login(c *cli.Context, mode string) *zabbix.API {
	var server string
	var user string
	var password string

	if server = os.Getenv(envZabbixURL); server == "" {
		server = c.GlobalString("server")
	}
	if user = os.Getenv(envZabbixUser); user == "" {
		user = c.GlobalString("user")
	}
	if password = os.Getenv(envZabbixPassword); password == "" {
		password = c.GlobalString("password")
	}

	logger.Log("info", "start Login() for "+mode+" cmd: "+
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
