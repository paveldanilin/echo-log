package event

type Event struct {
	fields map[string]string
}

func New() *Event {
	return &Event{
		fields: make(map[string]string),
	}
}

func (event *Event) HasField(name string) bool {
	_, ok := event.fields[name]
	return ok
}

func (event *Event) GetField(name string) string {
	if v, ok := event.fields[name]; ok {
		return v
	}
	return ""
}

func (event *Event) SetField(name string, value string) {
	event.fields[name] = value
}

func (event *Event) GetFields() map[string]string {
	return event.fields
}
