package json

import (
	"fmt"
	"testing"

	"github.com/paveldanilin/logwatch/internal/event"
)

func TestParser(t *testing.T) {

	eventString := `{"evenTtime": "2005-11-17 13:14:15", "level": "INFO", "message": "Hello message"}`

	eventDefinition := NewEventDefition()
	eventDefinition.SetField(NewFieldDefinition("eventtime", event.VALUE_DATETIME, "evenTtime"))
	eventDefinition.SetField(NewFieldDefinition("level", event.VALUE_STRING, "level"))
	eventDefinition.SetField(NewFieldDefinition("message", event.VALUE_STRING, "message"))

	p := NewParser(eventDefinition)

	evt, err := p.Parse(eventString)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v", evt.Map())
}
