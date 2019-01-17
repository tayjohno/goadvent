package main

import "testing"

func Example() {
	main()
	// Output:
	// 51833
	// 288
}

func TestPartOne(t *testing.T) {
	if PartOne() != 51833 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 51833)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 288 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 288)
	}
}
