package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
	"github.com/gefion-tech/gefion.gg/internal/util/gg"
)

func main() {
	res := gg.Root(os.Args[1:])
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
	fmt.Println(string(s))
}
