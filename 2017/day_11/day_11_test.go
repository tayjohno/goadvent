package main

import "testing"

func Example() {
	main()
	// Output:
	// 764
	// 1532
}

func TestPartOne(t *testing.T) {
	if PartOne() != 764 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 764)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 1532 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 1532)
	}
}
