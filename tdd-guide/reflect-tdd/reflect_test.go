package ref

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	s := struct {
		Name string
	}{expected}

	walk(s, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("expected %s got %s", expected, got[0])
	}
}
