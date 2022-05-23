package example01

import "testing"

func TestGreeting(t *testing.T) {
	var got = Greeting("George")
	want := "Hello George"
	if got != want {
		t.Fatalf("Expected %q", want, got)
	}
}
