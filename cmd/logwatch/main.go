package main

import (
	"fmt"

	"github.com/paveldanilin/logwatch/internal/event"
	"github.com/paveldanilin/logwatch/internal/file"
	"github.com/paveldanilin/logwatch/internal/script"
	"github.com/paveldanilin/logwatch/service"
)

func main() {
	ntf := service.NewNotifier()
	//ntf.Notify("bzz", "abc", "test")

	nameFieldDef := event.NewFieldDefinition("name", event.FIELD_STRING, map[string]interface{}{
		"csv_column_index": 0,
	})
	ageDef := event.NewFieldDefinition("age", event.FIELD_STRING, map[string]interface{}{
		"csv_column_index": 1,
	})
	emailDef := event.NewFieldDefinition("email", event.FIELD_STRING, map[string]interface{}{
		"csv_column_index": 2,
	})

	def := event.NewDefinition([]*event.FieldDefinition{
		nameFieldDef,
		ageDef,
		emailDef,
	})
	ep := event.NewCsvParser(def)
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
