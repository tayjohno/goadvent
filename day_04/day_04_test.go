package main

import "testing"

func Example() {
	main()
	// Output:
	// 477
	// 167
}

func TestPartOne(t *testing.T) {
	if PartOne() != 477 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 477)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 167 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 167)
	}
}
