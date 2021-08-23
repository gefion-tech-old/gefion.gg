package gg

import (
	"flag"
	"fmt"
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
func (command *TagCommandS) Run() error {
	switch command.fs.Arg(0) {

	case "clone":
		if len(command.fs.Args()) == 5 {
			fmt.Println("Скачивание с гита")
			return nil
		} else if len(command.fs.Args()) < 5 || len(command.fs.Args()) > 5 {
			return fmt.Errorf("Invalid number of arguments to run command `clone`")
		} else {
			return fmt.Errorf("Unknown error")
		}

	case "list":
		if len(command.fs.Args()) == 2 {
			fmt.Println("Получаю список тегов")
			return nil
		} else if len(command.fs.Args()) < 2 || len(command.fs.Args()) > 2 {
			return fmt.Errorf("Invalid number of arguments to run command `list`")
		} else {
			return fmt.Errorf("Unknown error")
		}

	default:
		return fmt.Errorf("Unknown argument")
	}
}
