package main

import "fmt"

var Version string

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print zabton version",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdVersion.Flag.BoolVar(&flagA, "a", false, "")
}

func runVersion(args []string) int {
	fmt.Printf("%s\n", Version)
	return 0
}
