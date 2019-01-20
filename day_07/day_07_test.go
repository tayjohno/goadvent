package main

import "testing"

func Example() {
	main()
	// Output:
	// azqje
	// -1
}

func TestPartOne(t *testing.T) {
	if PartOne() != "azqje" {
		t.Errorf("Answer was incorrect, got: %s, expected: %s.", PartOne(), "azqje")
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != -1 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), -1)
	}
}
