package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "zabton",
	Short: "Zabbix API and CLI tool set.",
	Long:  "zabton is a tool to make easy-to-use Zabbix by DevOps.",
}
