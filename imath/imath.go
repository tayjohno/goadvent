package imath

// Abs returns the absolute value of an int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sum returns the sum of a slice of integers
func Sum(integers []int) (sum int) {
	for i := 0; i < len(integers); i++ {
		sum += integers[i]
	}
	return
}

// Mod returns the value of (a % b)
func Mod(a int, b int) int {
	return a - (a / b * b)
}

// Min returns the smaller of two integers
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of two integers
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
