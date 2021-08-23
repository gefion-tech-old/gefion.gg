package gg

import (
	"flag"
	"fmt"
)

type TagCommandS struct {
	fs *flag.FlagSet
}

func TagCommand() *TagCommandS {
	tagCommand := &TagCommandS{
		fs: flag.NewFlagSet("tag", flag.ContinueOnError),
	}
	return tagCommand
}

func (command *TagCommandS) Name() string {
	return command.fs.Name()
}

func (command *TagCommandS) Init(args []string) error {
	return command.fs.Parse(args)
}

func (command *TagCommandS) Run() error {
	if command.fs.Arg(0) == "clone" && len(command.fs.Args()) == 5 {
		fmt.Println("Скачивание с гита")
		return nil
	} else if command.fs.Arg(0) == "clone" && len(command.fs.Args()) < 5 || len(command.fs.Args()) > 5 {
		return fmt.Errorf("Invalid number of arguments to run command `clone`")
	} else if command.fs.Arg(0) == "list" && len(command.fs.Args()) == 2 {
		fmt.Println("Получаю список тегов")
		return nil
	} else if command.fs.Arg(0) == "list" && len(command.fs.Args()) < 2 || len(command.fs.Args()) > 2 {
		return fmt.Errorf("Invalid number of arguments to run command `list`")
	} else {
		return fmt.Errorf("Unknown error")
	}
}
