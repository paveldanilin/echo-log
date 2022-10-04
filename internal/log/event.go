package log

import (
	"fmt"
)

type Event struct {
	fields map[string]string
}

func NewEvent() *Event {
	return &Event{
		fields: make(map[string]string),
	}
}

func (event *Event) HasField(name string) bool {
	_, ok := event.fields[name]
	return ok
}

func (le *Event) GetField(name string) (string, error) {
	if v, ok := le.fields[name]; ok {
		return v, nil
	}
	return "", fmt.Errorf("field noes not exists")
}

func (event *Event) SetField(name string, value string) {
	event.fields[name] = value
}
