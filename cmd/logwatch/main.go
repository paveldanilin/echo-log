package main

import (
	"fmt"

	"github.com/paveldanilin/logwatch/internal/event"
	"github.com/paveldanilin/logwatch/internal/file"
	"github.com/paveldanilin/logwatch/internal/script"
	"github.com/paveldanilin/logwatch/service"
)

func getCsvDefinition() *event.Definition {
	nameFieldDef := event.NewFieldDefinition("name", event.FIELD_STRING, map[string]interface{}{
		"csv.field_index": 0,
	})
	ageDef := event.NewFieldDefinition("age", event.FIELD_STRING, map[string]interface{}{
		"csv.field_index": 1,
	})
	emailDef := event.NewFieldDefinition("email", event.FIELD_STRING, map[string]interface{}{
		"csv.field_index": 2,
	})

	return event.NewDefinition([]*event.FieldDefinition{
		nameFieldDef,
		ageDef,
		emailDef,
	})
}

func getPatternDefinition() *event.Definition {
	dateTime := event.NewFieldDefinition("datetime", event.FIELD_STRING, map[string]interface{}{})
	logLevel := event.NewFieldDefinition("loglevel", event.FIELD_STRING, map[string]interface{}{})
	message := event.NewFieldDefinition("message", event.FIELD_STRING, map[string]interface{}{})

	return event.NewDefinition([]*event.FieldDefinition{
		dateTime,
		logLevel,
		message,
	})
}

func main() {
	ntf := service.NewNotifier()
	//ntf.Notify("bzz", "abc", "test")

	// def := getCsvDefinition()
	// ep := event.NewCsvParser(def, ";")

	def := getPatternDefinition()
	ep := event.NewPatternParser(def, `(?P<datetime>\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}) \[(?P<loglevel> {0,}(INFO|ERROR|DEBUG|WARNING))\] (?P<message>.*)$`)

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
