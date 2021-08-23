package main

import (
	"fmt"
	"os"

	_ "github.com/go-git/go-git"
)

func main() {
	if err := Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
