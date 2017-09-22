package zabbix

import (
	"testing"
)

func TestVersion(t *testing.T) {
	api := NewAPI(
		testZabbixURL,
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
		testZabbixURL,
		testZabbixUser,
		testZabbixPassword)

	auth, err := api.Login()

	if err != nil {
		t.Fatalf("Login() failed: %s", err)
	}

	t.Logf("authentication token: %s", auth)
}

func TestGetHost(t *testing.T) {
	testAPI := NewAPI(
		testZabbixURL,
		testZabbixUser,
		testZabbixPassword)
	testAPI.Auth = testAuth

	params := make(map[string]interface{})
	params["output"] = "extend"
	params["selectGroups"] = "extend"
	params["selectParentTemplates"] = "extend"

	hosts, err := testAPI.GetHost(params)

	if err != nil {
		t.Fatalf("GetHost() failed: %s", err)
	}

	found := false
	for _, obj := range hosts {
		o, ok := obj.(map[string]interface{})
		if ok {
			if v, ok := o["name"]; ok {
				found = true
				if v != "Zabbix server" {
					t.Errorf("GetHost() failed: %s", v)
				}
			}
			if v, ok := o["hostid"]; ok {
				found = true
				if v != "10084" {
					t.Errorf("GetHost() failed: %s", v)
				}
			}
		} else {
			t.Errorf("Irregular format: %T", obj)
		}
	}

	if !found {
		t.Error("GetHost() failed: empty")
	}
}
