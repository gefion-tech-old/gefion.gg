package gg

import (
	"flag"
	"fmt"

	"github.com/gefion-tech/gefion.gg/internal/util/model"
)

type TagCommandS struct {
	fs *flag.FlagSet
}

// Субкоманда `tag`
func TagCommand() *TagCommandS {
	tagCommand := &TagCommandS{
		fs: flag.NewFlagSet("tag", flag.ContinueOnError),
	}
	return tagCommand
}

// Получить имя субкоманды
func (command *TagCommandS) Name() string {
	return command.fs.Name()
}

// Получить инициализировать субкоманду
func (command *TagCommandS) Init(args []string) error {
	return command.fs.Parse(args)
}

// Получить выполнить субкоманду
func (command *TagCommandS) Run() model.Response {
	switch command.fs.Arg(0) {

	case "clone":
		if len(command.fs.Args()) == 5 {
			fmt.Println("Скачивание с гита")
			return model.CreateResponse("", "")
		} else if len(command.fs.Args()) < 5 || len(command.fs.Args()) > 5 {
			return model.CreateResponse("syntax error",
				"Invalid number of arguments to run command `clone`")
		} else {
			return model.CreateResponse("syntax error",
				"Unknown error")
		}

	case "list":
		if len(command.fs.Args()) == 2 {
			fmt.Println("Получаю список тегов")
			return model.CreateResponse("", "")
		} else if len(command.fs.Args()) < 2 || len(command.fs.Args()) > 2 {
			return model.CreateResponse("syntax error",
				"Invalid number of arguments to run command `list`")
		} else {
			return model.CreateResponse("syntax error",
				"Unknown error")
		}

	default:
		return model.CreateResponse("syntax error",
			"Unknown argument")

	}
}
