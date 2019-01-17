package main

import "testing"

func Example() {
	main()
	// Output:
	// 318883
	// 23948711
}

func TestPartOne(t *testing.T) {
	if PartOne() != 318883 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 318883)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 23948711 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 23948711)
	}
}
