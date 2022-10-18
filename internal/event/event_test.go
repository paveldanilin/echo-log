package event

import (
	"testing"
)

func TestEventSetGetField(t *testing.T) {
	e := New()
	e.SetValue("name", "Batman", VALUE_STRING)

	want := "Batman"
	got := e.Field("name").Raw()

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestEventHasField(t *testing.T) {
	e := New()
	e.SetValue("name", "Batman", VALUE_STRING)

	want := true
	got := e.Has("name")

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestEventDoesNotHaveField(t *testing.T) {
	e := New()
	e.SetValue("name", "Batman", VALUE_STRING)

	want := false
	got := e.Has("age")

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
