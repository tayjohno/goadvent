package two

import "testing"

// func TestLateralDistance(t *testing.T) {
// 	if LateralDistance(1) != 0 {
// 		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(1), 0)
// 	} else if LateralDistance(2) != 1 {
// 		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(2), 1)
// 	} else if LateralDistance(9) != 1 {
// 		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(9), 1)
// 	} else if LateralDistance(10) != 2 {
// 		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(10), 2)
// 	}
// }

func TestCountTo(t *testing.T) {
	if CountTo(1) != 2 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", CountTo(1), 2)
	} else if CountTo(20) != 23 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", CountTo(20), 23)
	} else if CountTo(100) != 122 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", CountTo(100), 122)
	}
}
