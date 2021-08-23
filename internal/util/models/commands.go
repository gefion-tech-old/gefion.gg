package models

import (
	"flag"
)

type GreetCommand struct {
	fs *flag.FlagSet

	name string
}
