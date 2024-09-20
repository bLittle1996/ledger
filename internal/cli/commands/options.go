package commands

import "github.com/spf13/pflag"

type CommandOption func(c *Command)

func DescOption(desc string) CommandOption {
	return CommandOption(func(c *Command) {
		c.Desc = desc
	})
}

func FlagStringOption(name, shorthand, defaultValue, desc string) CommandOption {
	return CommandOption(func(c *Command) {
		createCommandFlagSetIfNecessary(c)

		c.Flags.StringP(name, shorthand, defaultValue, desc)
	})
}

func SubCommandOption(subCmd *Command) CommandOption {
	return CommandOption(func(c *Command) {
		c.SubCommands = append(c.SubCommands, subCmd)
	})
}

func FlagBoolOption(name, shorthand string, defaultValue bool, desc string) CommandOption {
	return CommandOption(func(c *Command) {
		createCommandFlagSetIfNecessary(c)

		c.Flags.BoolP(name, shorthand, defaultValue, desc)
	})
}

func AllowUnknownFlagsOption(allowed bool) CommandOption {
	return CommandOption(func(c *Command) {
		createCommandFlagSetIfNecessary(c)

		c.Flags.ParseErrorsWhitelist.UnknownFlags = allowed
	})
}

func createCommandFlagSetIfNecessary(c *Command) {
	if c.Flags == nil {
		c.Flags = pflag.NewFlagSet(c.Name, pflag.ContinueOnError)
	}
}
