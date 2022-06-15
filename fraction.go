package hexgrid

import "math"

// FractionHex provides a more precise representation for hexagons when precision is required.
// It's also represented in Cube Coordinates
type FractionalHex struct {
	Q float64
	R float64
	S float64
}

func NewFractionalHex(q, r float64) FractionalHex {
	return FractionalHex{Q: q, R: r, S: -q - r}
}

// Round returns a 'rounded' FractionalHex, returning a Regular Hex
func (h FractionalHex) Round() Hex {
	roundToInt := func(a float64) int {
		if a < 0 {
			return int(a - 0.5)
		}
		return int(a + 0.5)
	}

	q := roundToInt(h.Q)
	r := roundToInt(h.R)
	s := roundToInt(h.S)

	qDiff := math.Abs(float64(q) - h.Q)
	rDiff := math.Abs(float64(r) - h.R)
	sDiff := math.Abs(float64(s) - h.S)

	if qDiff > rDiff && qDiff > sDiff {
		q = -r - s
	} else if rDiff > sDiff {
		r = -q - s
	} else {
		s = -q - r
	}
	return Hex{q, r, s}
}
