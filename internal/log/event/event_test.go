package event

import (
	"testing"
)

func TestEventSetGetField(t *testing.T) {
	e := New()
	e.SetField("name", "Batman")

	want := "Batman"
	got, err := e.GetField("name")

	if err != nil {
		t.Errorf("got error: %s", err)
	}

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestEventHasField(t *testing.T) {
	e := New()
	e.SetField("name", "Batman")

	want := true
	got := e.HasField("name")

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestEventDoesNotHaveField(t *testing.T) {
	e := New()
	e.SetField("name", "Batman")

	want := false
	got := e.HasField("age")

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

// --

func TestDef(t *testing.T) {
	f1 := NewFieldDefinition("name", FIELD_STRING)
	f2 := NewFieldDefinition("age", FIELD_NUMBER)

	ed := NewDefinition([]*FieldDefinition{f1, f2})

	ageDef, err := ed.GetFieldDefinition("age")
	if err != nil {
		panic(err)
	}

	if ageDef.GetFieldType() != FIELD_NUMBER {
		t.Error("zzz")
	}
}
