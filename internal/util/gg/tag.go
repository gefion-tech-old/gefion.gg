package gg

import (
	"flag"
	"fmt"
	"os"

	"github.com/gefion-tech/gefion.gg/internal/util/model"
	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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
func (command *TagCommandS) Run(u model.User) model.Response {
	switch command.fs.Arg(0) {
	case "clone":
		if len(command.fs.Args()) == 5 {
			Info("git tag clone " + command.fs.Arg(2))

			_, err := git.PlainClone(command.fs.Arg(3), false, &git.CloneOptions{
				URL: command.fs.Arg(2),
				Auth: &http.BasicAuth{
					Username: u.Username,
					Password: u.Password,
				},
				Progress: os.Stdout,
			})

			if err != nil {
				return model.CreateResponse(model.GIT__ERROR, err.Error())
			} else {
				return model.CreateResponse("", "")
			}

		} else if len(command.fs.Args()) < 5 || len(command.fs.Args()) > 5 {
			return model.CreateResponse(model.UTIL__ERROR,
				"Invalid number of arguments to run command `clone`")
		} else {
			return model.CreateResponse(model.UTIL__ERROR,
				model.UNDEFINED_UTIL__ERROR)
		}

	case "list":
		if len(command.fs.Args()) == 2 {
			fmt.Println("Получаю список тегов")
			return model.CreateResponse("", "")
		} else if len(command.fs.Args()) < 2 || len(command.fs.Args()) > 2 {
			return model.CreateResponse(model.UTIL__ERROR,
				"Invalid number of arguments to run command `list`")
		} else {
			return model.CreateResponse(model.UTIL__ERROR,
				model.UNDEFINED_UTIL__ERROR)
		}

	default:
		if len(command.fs.Args()) > 0 {
			return model.CreateResponse(model.UTIL__ERROR,
				"Unknown argument")
		} else {
			return model.CreateResponse(model.UTIL__ERROR,
				"You must pass an argument")
		}
	}
}
