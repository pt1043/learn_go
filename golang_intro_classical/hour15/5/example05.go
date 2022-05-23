package example04

import "testing"

type GreetingTest struct {
	name   string
	locala string
	want   string
}

var greetingTests = []GreetingTest{
	{"George", "en-US", "Hello George"},
	{"Chloe", "fr-FR", "Bonjour Chloe"},
	{"Giuseppe", "it-IT", "Ciao Guiseppe"},
}

func TestGreeting(t *testing.T) {
	for _, test := range greetingTests {
		got := Greeting(test.name, test.locale)
		if got != test.want {
			t.Errorf("Greeting(%s, %s) = %v; want %v", test.name, test.locale,
				actual, test.want)
		}
	}
}
