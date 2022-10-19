package json

import (
	"testing"

	"github.com/paveldanilin/logwatch/internal/event"
)

func TestParser(t *testing.T) {

	eventString := `{"evenTtime": "2020|08|05 13|14|15", "level": "INFO", "message": "Hello message"}`

	eventDefinition := NewEventDefition()
	eventDefinition.
		SetField(NewFieldDefinition("eventtime", event.VALUE_DATETIME, "evenTtime")).
		SetParameter("format", "Y|m|d H|i|s").
		SetParameter("tz", "Asia/Vladivostok")
	eventDefinition.SetField(NewFieldDefinition("level", event.VALUE_STRING, "level"))
	eventDefinition.SetField(NewFieldDefinition("message", event.VALUE_STRING, "message"))

	p := NewParser(eventDefinition)

	evt, err := p.Parse(eventString)

	if err != nil {
		t.Error(err)
		return
	}

	want := "2020-08-05 13:14:15 +1000 +10"
	got := evt.Field("eventtime").String()

	if want != got {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
