package gg

import (
	"fmt"
	"os"

	"github.com/gefion-tech/gefion.gg/internal/util/model"
)

// Интерфейс субкоманды
type Runner interface {
	Init([]string) error
	Run() model.Response
	Name() string
}

// Разделитель для параметров авторизации
var delimiter int = 0

// Определение интерфейса субкоманд
func Root(args []string) model.Response {
	if len(args) < 1 {
		return model.CreateResponse("syntax error", "You must pass a sub-command")
	}

	//Список подключенных субкоманд
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

	// Если передaна какая-то дичь
	return model.CreateResponse("syntax error",
		fmt.Errorf("Unknown subcommand: `%s`",
			subcommand).Error())
}
