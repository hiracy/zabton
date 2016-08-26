package main

import (
	"fmt"

	"github.com/hiracy/zabton/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
