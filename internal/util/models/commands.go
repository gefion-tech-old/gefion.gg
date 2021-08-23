package models

import (
	"flag"
)

type GreetCommand struct {
	Fs *flag.FlagSet

	Username string
}
