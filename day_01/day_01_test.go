package main

import "testing"

func Example() {
	main()
	// Output:
	// 995
	// 1130
}

func TestPartOne(t *testing.T) {
	if PartOne() != 995 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 995)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 1130 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 1130)
	}
}
