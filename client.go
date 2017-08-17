package main

import (
	"fmt"

	"github.com/hiracy/zabton/zabbix"
)

// Client is the client that interpose zabton and zabbix.
type Client struct {
	objects []string
	api     *zabbix.API
}

// NewClient creates new zabton client object.
func NewClient(object []string, api *zabbix.API) *Client {
	return &Client{object, api}
}

// PullHost download Host infomation from zabbix server.
func (client *Client) PullHost() error {
	fmt.Println("pull host")

	return nil
}

// PullHostGroup download Hostgroup infomation from zabbix server.
func (client *Client) PullHostgroup() error {
	fmt.Println("pull hostgroup")

	return nil
}
