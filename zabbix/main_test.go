package zabbix

import (
	"os"
	"testing"
)

var testZabbixVersion string
var testZabbixURL string
var testZabbixUser string
var testZabbixPassword string

const (
	zabbixAPIVersionDefault  = "3.2.7"
	zabbixAPIURLDefault      = "http://localhost:8080/api_jsonrpc.php"
	zabbixAPIUserDefault     = "Admin"
	zabbixAPIPasswordDefault = "zabbix"
)

func TestMain(m *testing.M) {
	initialize()
	exitCode := m.Run()
	finalize()
	os.Exit(exitCode)
}

func initialize() {
	if testZabbixVersion = os.Getenv("ZABBIX_API_VERSION"); testZabbixVersion == "" {
		testZabbixVersion = zabbixAPIVersionDefault
	}
	if testZabbixURL = os.Getenv("ZABBIX_API_URL"); testZabbixURL == "" {
		testZabbixURL = zabbixAPIURLDefault
	}
	if testZabbixUser = os.Getenv("ZABBIX_API_USER"); testZabbixUser == "" {
		testZabbixUser = zabbixAPIUserDefault
	}
	if testZabbixPassword = os.Getenv("ZABBIX_API_PASSWORD"); testZabbixPassword == "" {
		testZabbixPassword = zabbixAPIPasswordDefault
	}
}

func finalize() {
}
