package main

var cmdSub = &Command{
	Run:       runSub,
	UsageLine: "sub ",
	Short:     "Sub-command",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdSub.Flag.BoolVar(&flagA, "a", false, "")
}

// runSub executes sub command and return exit code.
func runSub(args []string) int {

	return 0
}
