package expense

import (
	"fmt"
	"github.com/spf13/pflag"
)

var (
	flagAddCategory string
)

// Run parses the arguments for the expense CLI command and executes the appropriate command
func RunAdd(args []string) error {
	fs := addFlagSet()

	if err := fs.Parse(args); err != nil {
		return err
	}

	argsNoFlags := fs.Args()

	if len(argsNoFlags) != 2 {
		// TODO print help msg
		return fmt.Errorf("incorrect arguments, expected <name> <price>")
	}

	return nil
}

func addFlagSet() *pflag.FlagSet {
	fs := pflag.NewFlagSet("add", pflag.ExitOnError)

	fs.StringVarP(&flagAddCategory, "category", "c", "Miscellaneous", "sets the category of the expense")

	return fs
}
