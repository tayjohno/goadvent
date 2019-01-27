package main

import "testing"

func Example() {
	main()
	// Output:
	// 19591
	// 62e2204d2ca4f4924f6e7a80f1288786
}

func TestPartOne(t *testing.T) {
	if PartOne() != 19591 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 19591)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != "62e2204d2ca4f4924f6e7a80f1288786" {
		t.Errorf("Answer was incorrect, got: %s, expected: %s.", PartTwo(), "62e2204d2ca4f4924f6e7a80f1288786")
	}
}
