package main

import (
	"fmt"
	"os"

	gg "github.com/gefion-tech/gefion.gg/internal/util"
	_ "github.com/go-git/go-git"
)

func main() {
	if err := gg.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
