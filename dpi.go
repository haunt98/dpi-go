package dpi

import "math"

// Unit is pixel
type Resolution struct {
	width  int64
	height int64
}

func NewResolution(width, height int64) Resolution {
	return Resolution{
		width:  width,
		height: height,
	}
}

func (r Resolution) CalculateDiagonal() float64 {
	return math.Sqrt(float64(r.width*r.width + r.height*r.height))
}

// Unit diagonal is inch
func CalculateDPI(r Resolution, diagonal float64) float64 {
	if diagonal == 0 {
		return 0
	}

	return r.CalculateDiagonal() / diagonal
}
