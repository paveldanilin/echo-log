package pattern

import (
	"testing"

	"github.com/paveldanilin/logwatch/internal/event"
)

func TestParser(t *testing.T) {

	eventDefiniton := NewEventDefition(`\[(?P<EventTime>.*)\] (?P<Method>GET|POST|PATCH|DELETE|HEAD) (?P<Path>/) (?P<Protocol>HTTP/1\.1) (?P<Status>\d{3}) (?P<ResponseSize>\d+)$`)
	eventDefiniton.SetField(NewFieldDefinition("eventtime", event.FIELD_DATETIME, "EventTime"))
	eventDefiniton.SetField(NewFieldDefinition("method", event.FIELD_STRING, "Method"))
	eventDefiniton.SetField(NewFieldDefinition("path", event.FIELD_STRING, "Path"))
	eventDefiniton.SetField(NewFieldDefinition("protocol", event.FIELD_STRING, "Protocol"))
	eventDefiniton.SetField(NewFieldDefinition("status", event.FIELD_STRING, "Status"))
	eventDefiniton.SetField(NewFieldDefinition("rs", event.FIELD_STRING, "ResponseSize"))

	definitions := make([]*EventDefinition, 0)
	definitions = append(definitions, eventDefiniton)

	parser := NewParser(definitions)

	evt, err := parser.Parse("[2018-03-19 22:10:18] GET / HTTP/1.1 200 1863")

	if err != nil {
		t.Error(err)
		return
	}

	wantTime := int64(1521461418)
	gotTime := evt.DatetimeValue("eventtime").Value().Unix()

	if wantTime != gotTime {
		t.Errorf("eventtime: got %d, wanted %d", gotTime, wantTime)
	}
}
