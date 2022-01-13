package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Jai")
	expect := "Hello, Jai\n"
	if got != expect {
		t.Fatalf("Expected %q but got %q", expect, got)
	}
}

func TestDepart(t *testing.T) {
	got := depart("Jai")
	expect := "Bye, Jai\n"
	if got != expect {
		t.Fatalf("Expected %q but got %q", expect, got)
	}
}

func TestTableDrivenGreet(t *testing.T) {
	scenarios := []struct {
		input  string
		output string
	}{
		{input: "Gopher", output: "Hello, Gopher\n"},
		{input: "", output: "Hello, \n"},
	}

	for _, scenario := range scenarios {
		got := Greet(scenario.input)
		if got != scenario.output {
			t.Errorf("For scenario '%+v', expceted value was %q but found %q", scenario, scenario.output, got)
		}
	}
}
