package main

import "testing"

func Example() {
	main()
	// Output:
	// 0
	// 0
}

func TestPartOne(t *testing.T) {
	if PartOne() != 0 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 0)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 0 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 0)
	}
}
