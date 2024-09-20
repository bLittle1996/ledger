package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

type Command struct {
	// The name of the command, must not contain spaces
	Name string
	// A brief description of the command
	Desc string
	// Flags that modify the commands behaviour
	Flags *pflag.FlagSet
	// Commands that can be accessed from this command.
	// Subcommands are ran via `CommandName SubCommand`.
	// If a subcommand is ran, the parent command is not executed.
	SubCommands []*Command
	// Execute is the code that runs when the command is executed!
	Execute func(c *Command, args []string) error
}

func NewCommand(name string, handler func(c *Command, args []string) error, options ...CommandOption) *Command {
	if strings.Contains(name, " ") {
		fmt.Fprintf(os.Stderr, "command must not have spaces. received: %s\n", name)
		os.Exit(1)
		return nil
	}

	cmd := &Command{
		Name:    name,
		Execute: handler,
	}

	for _, opt := range options {
		opt(cmd)
	}

	return cmd
}

// exec is a wrapper around Execute that parses the arguments before calling Execute
func (c *Command) exec(args []string) error {
	remainingArgs, err := c.parseFlags(args)

	if err != nil {
		return err
	}

	return c.Execute(c, remainingArgs)
}

// Run determines if this command should be executed or if a subcommand should be executed based on the arguments provided.
func (c *Command) Run(args []string) error {
	if len(args) == 0 || len(c.SubCommands) == 0 {
		return c.exec(args)
	}

	firstNonFlagArg := args[0]
	var matchingSubCmd *Command

	for _, subCmd := range c.SubCommands {
		if subCmd.Name == firstNonFlagArg {
			matchingSubCmd = subCmd
			break
		}
	}

	// Note we _run_ the sub command so that it handles all the arg parsing
	// and execution of its  own potential sub commands
	if matchingSubCmd != nil {
		return matchingSubCmd.Run(args[1:]) // skip the first arg, since it is the command name
	}

	return c.exec(args)
}

// parseFlags parses the flags in the command's FlagSet and returns the remaining non-flag args
// The command name cannot be the first argument when parsing.
func (c *Command) parseFlags(args []string) ([]string, error) {
	if c.Flags != nil {
		err := c.Flags.Parse(args)
		return c.Flags.Args(), err
	}

	return []string{}, nil
}
