package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jo")

	got := buffer.String()
	want := "Hello, Jo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
