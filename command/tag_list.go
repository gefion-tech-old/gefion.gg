package command

import (
	"flag"
	"fmt"

	m "github.com/gefion-tech/gefion.gg/model"
)

type TagList struct {
	fs *flag.FlagSet
	remote,
	username,
	password string
}

func TagListCommand() *TagList {
	tl := &TagList{
		fs: flag.NewFlagSet("tag:list", flag.ContinueOnError),
	}

	tl.fs.StringVar(&tl.username, "remote", "", "Remote host.")
	tl.fs.StringVar(&tl.username, "username", "", "Git uses a username to associate commits with an identity.")
	tl.fs.StringVar(&tl.password, "password", "", "Your account password or token.")

	return tl
}

func (g *TagList) Name() string {
	return g.fs.Name()
}

func (g *TagList) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *TagList) Run() m.Error {
	fmt.Println("Hello", g.username, "!")
	return m.MakeRes(nil)
}
