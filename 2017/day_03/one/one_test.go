package one

import "testing"

func TestLateralDistance(t *testing.T) {
	if LateralDistance(1) != 0 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(1), 0)
	} else if LateralDistance(2) != 1 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(2), 1)
	} else if LateralDistance(9) != 1 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(9), 1)
	} else if LateralDistance(10) != 2 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", LateralDistance(10), 2)
	}
}

func TestRotationalDistancer(t *testing.T) {
	if RotationalDistance(1) != 0 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", RotationalDistance(1), 0)
	} else if RotationalDistance(2) != 0 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", RotationalDistance(2), 0)
	} else if RotationalDistance(9) != 1 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", RotationalDistance(9), 1)
	} else if RotationalDistance(10) != 1 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", RotationalDistance(10), 1)
	} else if RotationalDistance(46) != 0 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", RotationalDistance(46), 0)
	} else if RotationalDistance(49) != 3 {
		t.Errorf("Answer was incorrect, got: %d, expected: %d.", RotationalDistance(49), 3)
	}
}
