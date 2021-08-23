package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/gefion-tech/gefion.gg/internal/util/gg"
)

func main() {
	if err := gg.Root(os.Args[1:]); err != nil {
		color.Red("Error: " + err.Error())
		os.Exit(1)
	}
}
