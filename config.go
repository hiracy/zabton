package main

import (
	"github.com/BurntSushi/toml"
)

// Config is the general application config
type Config struct {
	Zabbix ZabbixConfiguration
}

// ZabbixConfiguration is the config by Zabbix
type ZabbixConfiguration struct {
	Editable EditableConfiguration
}

// EditableConfiguration is the config by Zabbix editable objects
type EditableConfiguration struct {
	Host []string `toml:"host"`
}

// LoadConfig loads a Config from a file.
func LoadConfig(path string) (editables *EditableConfiguration, err error) {
	var config Config
	_, err = toml.DecodeFile(path, &config)

	if err != nil {
		return nil, err
	}

	return &config.Zabbix.Editable, nil
}
