package dpi

// GCD is greatest common divisor
func calculateGCD(a, b int64) int64 {
	if a < b {
		a, b = b, a
	}

	for b != 0 {
		a = a % b
		a, b = b, a
	}

	return a
}
