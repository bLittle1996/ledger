package expense

import "fmt"

// Run parses the arguments for the expense CLI command and executes the appropriate command
func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("unable to run expense command, no arguments provided")
	}
	cmd := args[0]
	cmdArgs := args[1:]

	switch cmd {
	case "add":
		return RunAdd(cmdArgs)
	default:
		return fmt.Errorf("unrecognized expense command: %s", cmd)
	}
}
