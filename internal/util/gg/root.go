package gg

import (
	"errors"
	"fmt"
	"os"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

// Определение интерфейса субкоманд []Runner
func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		TagCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}
