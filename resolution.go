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

func (r Resolution) CalculateAspectRatio() (int64, int64) {
	if r.width == 0 || r.height == 0 {
		return 0, 0
	}

	gcd := calculateGCD(r.width, r.height)

	return r.width / gcd, r.height / gcd
}
