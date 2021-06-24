package resolution

// PPI is pixels per inch
// Unit diagonal is inch
func PPI(r Resolution, diagonal float64) float64 {
	if diagonal == 0 {
		return 0
	}

	return r.Diagonal() / diagonal
}
