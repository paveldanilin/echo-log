package main

import (
	"fmt"

	"github.com/paveldanilin/logwatch/internal/file"
	"github.com/paveldanilin/logwatch/internal/script"
	"github.com/paveldanilin/logwatch/service"
)

func main() {
	ntf := service.NewNotifier()
	//ntf.Notify("bzz", "abc", "test")

	s := script.NewLuaScript()
	s.LoadFile("filter.lua")

	s.Register("ntf", ntf)

	w := file.NewTailWatcher(&file.TailWatcherConfig{
		Filename: "./out.log",
		// Follow:   true,
	})
	w.Watch(func(line string) {
		// parse
		e := ep.Parse(line)
		if e == nil {
			//fmt.Println("=>>>>" + line)
			return
		}

		// TODO: decode line to event
		r, err := s.Call("lw_filter_event", e)
		if err != nil {
			panic(err)
		}
		if r.(bool) {
			fmt.Println(">>" + line)
		}
	})
}
