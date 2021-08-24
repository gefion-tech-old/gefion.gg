package gg

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gefion-tech/gefion.gg/internal/util/model"
)

type User struct {
	username string
	password string
}

// Интерфейс субкоманды
type Runner interface {
	Init([]string) error
	Run() model.Response
	Name() string
}

// Разделитель для параметров авторизации
var delimiter int = 0

var user User

// Определение интерфейса субкоманд
func Root(args []string) model.Response {
	if len(args) < 1 {
		return model.CreateResponse(model.UTIL__ERROR, "You must pass a subcommand")
	}

	// Устанавливаю параметры для авторизации
	user.username = model.GetFlagValue(args, model.Usrn_s, model.Usrn_f, "")
	user.password = model.GetFlagValue(args, model.Pass_s, model.Pass_f, "")

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
		// Если найдена поддерживаемая субкоманда
		if cmd.Name() == subcommand {
			// Если разделитель присутствует и после него указано
			// минимальное число парматеров.
			if delimiter > 0 && len(os.Args[delimiter:]) > 2 {
				color.Yellow("Working with a private repository...")
				cmd.Init(os.Args[2:delimiter])
			} else {
				color.Yellow("Working with a public repository...")
				// Если разделитель присутствует но после него
				// не указаны параметры.
				if delimiter > 0 && len(os.Args[delimiter:]) > 0 {
					cmd.Init(os.Args[2:delimiter])
				} else {
					cmd.Init(os.Args[2:])
				}
			}
			return cmd.Run()
		}
	}

	// Если передaна какая-то дичь
	return model.CreateResponse(model.UTIL__ERROR,
		fmt.Errorf("Unknown subcommand: `%s`",
			subcommand).Error())
}
