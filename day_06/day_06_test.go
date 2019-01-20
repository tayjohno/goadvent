package main

import "testing"

func Example() {
	main()
	// Output:
	// 11137
	// 1037
}

func TestPartOne(t *testing.T) {
	if PartOne() != 11137 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 11137)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 1037 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 1037)
	}
}
