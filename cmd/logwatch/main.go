package main

import (
	"fmt"
	"strings"

	"github.com/paveldanilin/logwatch/internal/log/event"
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
	ep := CsvEventParser{}
	s := script.NewLuaScript()
	s.LoadFile("filter.lua")

	w := service.NewTailFileWatcher("./test.log")
	w.Watch(func(line string) {

		// parse
		e := ep.Parse(line)

		// TODO: decode line to event
		r, err := s.Call("filter_event", e)
		if err != nil {
			panic(err)
		}
		if r.(bool) {
			fmt.Println(">>" + line)
		}
	})
}
