package main

import (
	"os"

	c "github.com/gefion-tech/gefion.gg/command"
	m "github.com/gefion-tech/gefion.gg/model"
)

type Runner interface {
	Init([]string) error
	Run() interface{}
	Name() string
}

func root(args []string) interface{} {
	if len(args) < 1 {
		return m.Error{
			Error: &m.ErrorBody{
				Type:    m.UTIL__ERROR,
				Message: "You must pass a subcommand",
			},
		}
	}

	subcommand := os.Args[1]
	cmds := []Runner{
		c.TagListCommand(),
		// c.TagCloneCommand(),
	}

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return m.Error{
		Error: &m.ErrorBody{
			Type:    m.UTIL__ERROR,
			Message: "Unknown subcommand: " + subcommand,
		},
	}
}
