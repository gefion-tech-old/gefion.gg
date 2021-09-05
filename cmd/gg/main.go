package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
)

func main() {
	res := root(os.Args[1:])
	json_data, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Форматированный вывод в консоль объекта ответа
	var obj map[string]interface{}
	json.Unmarshal([]byte(json_data), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 2
	s, _ := f.Marshal(obj)

	fmt.Printf("\n\n")
	fmt.Println(string(s))
}
