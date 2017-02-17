package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// Version is the package version
var Version string

func printVersion(c *cli.Context) {
	fmt.Fprintf(c.App.Writer, "%v\n", c.App.Version)
}
