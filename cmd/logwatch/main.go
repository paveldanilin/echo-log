package main

import (
	"fmt"
	"strings"

	"github.com/paveldanilin/logwatch/internal/event"
	"github.com/paveldanilin/logwatch/internal/file"
	"github.com/paveldanilin/logwatch/internal/script"
	"github.com/paveldanilin/logwatch/service"
)

type EventParser interface {
	Parse(text string) *event.Event
}

type CsvEventParser struct {
}

func (csv *CsvEventParser) Parse(text string) *event.Event {
	fields := strings.Split(text, ";")

	e := event.New()
	e.SetField("name", fields[0])
	e.SetField("age", fields[1])
	e.SetField("email", fields[2])

	return e
}

func main() {
	ntf := service.NewNotifier()
	//ntf.Notify("bzz", "abc", "test")

	ep := CsvEventParser{}
	s := script.NewLuaScript()
	s.LoadFile("filter.lua")

	s.Register("ntf", ntf)

	w := file.NewTailWatcher("./test.log")
	w.Watch(func(line string) {

		// parse
		e := ep.Parse(line)

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
