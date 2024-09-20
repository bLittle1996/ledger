package cli

import (
	"fmt"

	"github.com/bLittle1996/ledger/internal/cli/commands"
)

var rootCmd = commands.NewCommand(
	"ledger",
	func(c *commands.Command, args []string) error {
		if enabled, _ := c.Flags.GetBool("verbose"); enabled {
			fmt.Println("Verbose logging enabled...")
		}
		fmt.Println("EXECUTING", c.Name, args)

		return nil
	},
	commands.DescOption("The entrypoint for the ledger cli. Doesn't do anything on it's own!"),
	commands.AllowUnknownFlagsOption(true),
	commands.FlagBoolOption("verbose", "v", false, "enables additional logging when executing commands"),
	commands.SubCommandOption(
		commands.NewCommand(
			"expense",
			func(c *commands.Command, args []string) error {
				category, err := c.Flags.GetString("category")

				if err != nil {
					return err
				}

				fmt.Println("EXECUTING", c.Name, args)
				if enabled, _ := c.Flags.GetBool("verbose"); enabled {
					fmt.Println("SAVING EXPENSE TO DB, CATEGORY: ", category)
				}

				return nil
			},
			commands.FlagStringOption("category", "c", "Miscellaneous", "the category of the expense, such as groceries or bills"),
		),
	),
)

// Run runs the CLI application, parsing the arguments to determine what to do.
func Run(args []string) error {
	fmt.Println("ARGUMENTS", args)

	return rootCmd.Run(args)
}
