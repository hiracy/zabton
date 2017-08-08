package zabbix

import (
	"os"
	"testing"
)

var testZabbixVersion string
var testZabbixUrl string
var testZabbixUser string
var testZabbixPassword string

const (
	ZABBIX_API_VERSION_DEFAULT  = "3.2.7"
	ZABBIX_API_URL_DEFAULT      = "http://localhost:8080/api_jsonrpc.php"
	ZABBIX_API_USER_DEFAULT     = "Admin"
	ZABBIX_API_PASSWORD_DEFAULT = "zabbix"
)

func TestMain(m *testing.M) {
	initialize()
	exitCode := m.Run()
	finalize()
	os.Exit(exitCode)
}

func initialize() {
	if testZabbixVersion = os.Getenv("ZABBIX_API_VERSION"); testZabbixVersion == "" {
		testZabbixVersion = ZABBIX_API_VERSION_DEFAULT
	}
	if testZabbixUrl = os.Getenv("ZABBIX_API_URL"); testZabbixUrl == "" {
		testZabbixUrl = ZABBIX_API_URL_DEFAULT
	}
	if testZabbixUser = os.Getenv("ZABBIX_API_USER"); testZabbixUser == "" {
		testZabbixUser = ZABBIX_API_USER_DEFAULT
	}
	if testZabbixPassword = os.Getenv("ZABBIX_API_PASSWORD"); testZabbixPassword == "" {
		testZabbixPassword = ZABBIX_API_PASSWORD_DEFAULT
	}
}

func finalize() {
}
