package day

import "testing"

func Example() {
	main()
	// Output:
	// 552
	// 330785
}

func TestPartOne(t *testing.T) {
	if PartOne() != 552 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartOne(), 0)
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo() != 330785 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", PartTwo(), 0)
	}
}
