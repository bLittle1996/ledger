package cli

import (
	"fmt"

	"github.com/bLittle1996/ledger/internal/cli/expense"
)

// Run runs the CLI application, parsing the arguments to determine what to do.
func Run(args []string) error {
	cmd := args[0]
	subCmds := args[1:]

	switch cmd {
	case "expense":
		return expense.Run(subCmds)
	default:
		return fmt.Errorf("Unrecognized command: %s", cmd)
	}
}
