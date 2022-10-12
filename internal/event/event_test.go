package event

import (
	"testing"
)

func TestEventSetGetField(t *testing.T) {
	e := New()
	e.SetField("name", "Batman")

	want := "Batman"
	got := e.GetField("name")

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
