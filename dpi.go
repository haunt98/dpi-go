package dpi

// Unit diagonal is inch
func CalculateDPI(r Resolution, diagonal float64) float64 {
	if diagonal == 0 {
		return 0
	}

	return r.CalculateDiagonal() / diagonal
}
