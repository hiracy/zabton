package zabbix

import (
	"testing"
)

func TestVersion(t *testing.T) {
	api := NewAPI(
		testZabbixUrl,
		"",
		"")

	version, err := api.Version()

	if err != nil {
		t.Fatalf("Version() failed: %s", err)
	}

	if testZabbixVersion != version {
		t.Fatalf("Zabbix Server API Version: version should be %s, but %s", testZabbixVersion, version)
	}

	t.Logf("zabbix server api version: %s", version)
}

func TestLogin(t *testing.T) {
	api := NewAPI(
		testZabbixUrl,
		testZabbixUser,
		testZabbixPassword)

	auth, err := api.Login()

	if err != nil {
		t.Fatalf("Login() failed: %s", err)
	}

	t.Logf("authentication token: %s", auth)
}
