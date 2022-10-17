package csv

import (
	"testing"

	"github.com/paveldanilin/logwatch/internal/event"
)

func TestParser(t *testing.T) {
	name := NewFieldDefinition("name", event.FIELD_STRING, 0)
	age := NewFieldDefinition("age", event.FIELD_INT, 1)
	email := NewFieldDefinition("email", event.FIELD_STRING, 2)
	eventtime := NewFieldDefinition("eventtime", event.FIELD_DATETIME, 3)

	eventDefiniton := NewEventDefition()
	eventDefiniton.SetField(name)
	eventDefiniton.SetField(age)
	eventDefiniton.SetField(email)
	eventDefiniton.SetField(eventtime)

	parser := NewParser(eventDefiniton, ";")

	evt, err := parser.Parse("Jhon;54;jhon@mail.com;2020-08-05 13:14:15")

	if err != nil {
		t.Error(err)
	}

	wantName := "Jhon"
	wantAge := 54
	wantEmail := "jhon@mail.com"
	wantTime := int64(1596597255)

	gotName := evt.Field("name").Raw()
	gotAge := evt.Field("age").(event.IntValue).Value()
	gotEmail := evt.StringValue("email").Value()
	gotTime := evt.DatetimeValue("eventtime").Value().Unix()

	if wantName != gotName {
		t.Errorf("name: got %s, wanted %s", gotName, wantName)
	}

	if wantAge != gotAge {
		t.Errorf("age: got %d, wanted %d", gotAge, wantAge)
	}

	if wantEmail != gotEmail {
		t.Errorf("email: got %s, wanted %s", gotEmail, wantEmail)
	}

	if wantTime != gotTime {
		t.Errorf("email: got %d, wanted %d", gotTime, wantTime)
	}
}
