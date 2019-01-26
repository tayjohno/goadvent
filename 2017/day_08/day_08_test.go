package main

import "testing"

func Example() {
	main()
	// Output:
	// -1
	// -1
}

func TestPartOne(t *testing.T) {
	if PartOne() != -1 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), -1)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != -1 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), -1)
	}
}
