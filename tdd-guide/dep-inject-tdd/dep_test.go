package dep

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "chris")

	got := buffer.String()

	want := "Hello, chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
