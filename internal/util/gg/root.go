package gg

import (
	"errors"
	"fmt"
	"os"
)

// Интерфейс субкоманды
type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

// разделитель параметров авторизации
var delimiter int = 0

// Определение интерфейса субкоманд
func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	// Список подключенных субкоманд
	cmds := []Runner{
		TagCommand(),
	}

	subcommand := os.Args[1]

	// Поиск аргумента разделителя
	for i, cmd := range os.Args {
		if cmd == "--" {
			delimiter = i
		}
	}

	// Определение субкоманды
	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			if delimiter > 0 {
				cmd.Init(os.Args[2:delimiter])
			} else {
				cmd.Init(os.Args[2:])
			}
			return cmd.Run()
		}
	}

	// Если передена какая-то дичь
	return fmt.Errorf("Unknown subcommand: `%s`", subcommand)
}
