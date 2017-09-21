package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hiracy/zabton/logger"
	"github.com/hiracy/zabton/zabbix"
)

const (
	zabtonWritePermission = 0666
	zabtonJSONFilePrefix  = ""
	zabtonJSONFileIndent  = "    "
)

// Client is the client that interpose zabton and zabbix.
type Client struct {
	api       *zabbix.API
	writePath string
	readPath  string
	editables *EditableConfiguration
}

// NewClient creates new zabton client object.
func NewClient(api *zabbix.API, writepath, readpath string, editables *EditableConfiguration) *Client {
	return &Client{api, writepath, readpath, editables}
}

// PullHost download Host infomation from zabbix server.
func (client *Client) PullHost() error {
	logger.Log("info", "start PullHost()")

	params := make(map[string]interface{})
	params["output"] = "extend"
	params["selectGroups"] = "extend"
	params["selectParentTemplates"] = "extend"

	hosts, err := client.api.GetHost(params)

	if err != nil {
		return err
	}

	logger.Log("info", "start readAllZabbixObjects()")

	existingObjects, err := readAllZabbixObjects(client.writePath)

	if err != nil {
		return err
	}

	err = saveZabbixObjects(existingObjects, hosts, client.editables.Host, "host", client.writePath)

	if err != nil {
		return err
	}

	logger.Log("info", "succeeded PullHost()")

	return nil
}

// PullHostgroup download Hostgroup infomation from zabbix server.
func (client *Client) PullHostgroup() error {
	logger.Log("info", "start PullHostgroup(path="+client.writePath+")")
	return nil
}

func saveZabbixObjects(existingObjects map[string]interface{}, updateObjects []interface{}, editables []string, objectName, path string) error {
	var saving []interface{}

	for _, obj := range updateObjects {
		o, ok := obj.(map[string]interface{})
		if ok {
			content := map[string]interface{}{}
			for _, e := range editables {
				if v, ok := o[e]; ok {
					content[e] = v
				}
			}

			saving = append(saving, content)
		} else {
			return fmt.Errorf("Irregular format: %T", obj)
		}
	}

	existingObjects[objectName] = saving

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, zabtonWritePermission)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent(zabtonJSONFilePrefix, zabtonJSONFileIndent)
	err = encoder.Encode(existingObjects)
	if err != nil {
		return err
	}

	return nil
}

func readAllZabbixObjects(path string) (objects map[string]interface{}, err error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]interface{}), nil
		}
		return nil, err
	}
	defer f.Close()

	var obj interface{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&obj)
	if err != nil {
		return nil, err
	}

	if v, ok := obj.(map[string]interface{}); ok {
		return v, nil
	}

	return nil, fmt.Errorf("Irregular format: %s", path)
}
