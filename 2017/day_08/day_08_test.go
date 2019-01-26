package main

import "testing"

func Example() {
	main()
	// Output:
	// 5143
	// 6209
}

func TestPartOne(t *testing.T) {
	if PartOne() != 5143 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 5143)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 6209 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 6209)
	}
}
