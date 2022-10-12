package event

import (
	"testing"
)

func TestFieldDefinition(t *testing.T) {
	fd := NewFieldDefinition("name", FIELD_NUMBER)
	fd.SetParam("index", 123)

	want := 123
	got := fd.GetIntParam("index")

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
