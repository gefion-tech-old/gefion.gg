package command

import (
	"flag"
	"fmt"

	m "github.com/gefion-tech/gefion.gg/model"
)

type TagClone struct {
	fs *flag.FlagSet
	tag,
	remote,
	destination,
	hash,
	username,
	password string
}

func TagCloneCommand() *TagClone {
	tc := &TagClone{
		fs: flag.NewFlagSet("tag:clone", flag.ContinueOnError),
	}

	tc.fs.StringVar(&tc.tag, "t", "", "Tag.")
	tc.fs.StringVar(&tc.remote, "r", "", "Remote host.")
	tc.fs.StringVar(&tc.destination, "d", "", "Destination.")
	tc.fs.StringVar(&tc.hash, "h", "", "Hash.")
	tc.fs.StringVar(&tc.username, "username", "", "Git uses a username to associate commits with an identity.")
	tc.fs.StringVar(&tc.password, "password", "", "Your account password or token.")

	return tc
}

// Получить имя субкоманды
func (c *TagClone) Name() string {
	return c.fs.Name()
}

// Получить инициализировать субкоманду
func (g *TagClone) Init(args []string) error {
	return g.fs.Parse(args)
}

// Выполнить субкоманду
func (c *TagClone) Run() m.Error {
	fmt.Println("clone " + c.username)
	return m.MakeRes(nil)
}
