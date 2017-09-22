package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	editables, err := LoadConfig("zabton.conf")

	if err != nil {
		t.Fatalf("LoadConfig() failed: %s", err)
	}

	found := false
	var errStr []string

	for _, v := range editables.Host {
		switch v {
		case "name", "status":
			found = true
		default:
			found = false
			errStr = append(errStr, v)
			break
		}
	}

	if !found {
		t.Errorf("LoadConfig() failed: %v", errStr)
	}
}
