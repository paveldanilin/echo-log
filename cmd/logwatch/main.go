package main

import (
	"github.com/paveldanilin/logwatch/internal/script"
)

func main() {

	s := script.NewLuaScript()
	s.LoadString(`
		function zoo()
			return "zeebra"
		end
	`)

	ret, err := s.CallFunction("zoo")
	if err != nil {
		panic(err)
	}
	println(ret.(string))

}
