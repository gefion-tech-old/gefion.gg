package gg

import (
	"flag"
	"fmt"

	model "github.com/gefion-tech/gefion.gg/internal/util/models"
)

func NewGreetCommand() *model.GreetCommand {
	gc := model.GreetCommand{
		Fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}

	gc.Fs.StringVar(&gc.Username, "name", "World", "name of the person to be greeted")

	return gc
}

func (g *model.GreetCommand) Name() string {
	return g.Fs.Username()
}

func (g *model.GreetCommand) Init(args []string) error {
	return g.Fs.Parse(args)
}

func (g *model.GreetCommand) Run() error {
	fmt.Println("Hello", g.Username, "!")
	return nil
}
