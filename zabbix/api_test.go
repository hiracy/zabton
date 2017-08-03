package zabbix

import (
	"testing"
)

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
