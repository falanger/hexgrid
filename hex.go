package hexgrid

import (
	"fmt"
	"math"
)

var directions = []Hex{
	NewHex(1, 0),
	NewHex(1, -1),
	NewHex(0, -1),
	NewHex(-1, 0),
	NewHex(-1, +1),
	NewHex(0, +1),
}

// Hex describes a regular hexagon with Cube Coordinates (although the
// S coordinate is computed on the constructor)
//
// It's also easy to reference them as axial (trapezoidal
// coordinates):
// - R represents the vertical axis
// - Q the diagonal one
// - S can be ignored
//
// For additional reference on these coordinate systems, see
// http://www.redblobgames.com/grids/hexagons/#coordinates
//           _ _
//         /     \
//    _ _ /(0,-1) \ _ _
//  /     \  -R   /     \
// /(-1,0) \ _ _ /(1,-1) \
// \  -Q   /     \       /
//  \ _ _ / (0,0) \ _ _ /
//  /     \       /     \
// /(-1,1) \ _ _ / (1,0) \
// \       /     \  +Q   /
//  \ _ _ / (0,1) \ _ _ /
//        \  +R   /
//         \ _ _ /
type Hex struct {
	Q int // x axis
	R int // y axis
	S int // z axis
}

// NewHex constructs new Hex value with specified q and r.
func NewHex(q, r int) Hex {
	return Hex{Q: q, R: r, S: -q - r}
}

func (h Hex) String() string {
	return fmt.Sprintf("(%d,%d)", h.Q, h.R)
}

// Add performs an addition operation by adding the Hex 'b' to this Hex
func (h Hex) Add(b Hex) Hex {
	return NewHex(h.Q+b.Q, h.R+b.R)
}

// Subtract performas a subtraction operation by subtracting the Hex 'b' from this Hex
func (h Hex) Subtract(b Hex) Hex {
	return NewHex(h.Q-b.Q, h.R-b.R)
}

// Scale applies a scaling factor 'k' to the hex, returning a new
// Hex. If factor k is 1 there's no change
func (h Hex) Scale(k int) Hex {
	return NewHex(h.Q*k, h.R*k)
}

// Length calculates the length of a hex
func (h Hex) Length() int {
	return int((math.Abs(float64(h.Q)) + math.Abs(float64(h.R)) + math.Abs(float64(h.S))) / 2.)
}

// DistanceTo returns the distance from the Hex 'h' to the Hex 'b'
func (h Hex) DistanceTo(b Hex) int {
	sub := h.Subtract(b)
	return sub.Length()
}

// Neighbor returns the neighbor at a given direction
func (h Hex) Neighbor(dir Direction) Hex {
	directionOffset := directions[dir]
	return NewHex(h.Q+directionOffset.Q, h.R+directionOffset.R)
}

// LineDraw returns the slice of hexagons that exist on a line that
// goes from hexagon a to hexagon b
func (h Hex) LineDraw(b Hex) []Hex {
	hexLerp := func(a FractionalHex, b FractionalHex, t float64) FractionalHex {
		return NewFractionalHex(a.Q*(1-t)+b.Q*t, a.R*(1-t)+b.R*t)
	}

	N := h.DistanceTo(b)

	// Sometimes the hexLerp will output a point that’s on an edge.
	// On some systems, the rounding code will push that to one side or the other,
	// somewhat unpredictably and inconsistently.
	// To make it always push these points in the same direction, add an “epsilon” value to a.
	// This will “nudge” things in the same direction when it’s on an edge, and leave other points unaffected.
	aNudge := NewFractionalHex(float64(h.Q)+0.000001, float64(h.R)+0.000001)
	bNudge := NewFractionalHex(float64(b.Q)+0.000001, float64(b.R)+0.000001)

	results := make([]Hex, 0)
	step := 1. / math.Max(float64(N), 1)

	for i := 0; i <= N; i++ {
		results = append(results, hexLerp(aNudge, bNudge, step*float64(i)).Round())
	}
	return results
}

// Range returns the set of hexagons around this hex for a given
// radius
func (h Hex) Range(radius int) []Hex {
	return Range(h, radius)
}

// Range returns the set of hexagons around the center Hex for a given
// radius
func Range(center Hex, radius int) []Hex {
	var results []Hex

	if radius >= 0 {
		for dx := -radius; dx <= radius; dx++ {
			dy := math.Max(float64(-radius), float64(-dx-radius))
			for ; dy <= math.Min(float64(radius), float64(-dx+radius)); dy++ {
				results = append(results, center.Add(NewHex(int(dx), int(dy))))
			}
		}
	}

	return results
}

// HasLineOfSight determines if a given hexagon is visible from
// another hexagon, taking into consideration a set of blocking
// hexagons
func (h Hex) HasLineOfSight(target Hex, blocking []Hex) bool {
	contains := func(s []Hex, e Hex) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}

	for _, hh := range h.LineDraw(target) {
		if contains(blocking, hh) {
			return false
		}
	}

	return true
}

// FieldOfView feturns the list of hexagons that are visible from this Hex
func (h Hex) FieldOfView(candidates []Hex, blocking []Hex) []Hex {
	var results []Hex

	for _, c := range candidates {
		distance := h.DistanceTo(c)

		if len(blocking) == 0 || distance <= 1 || h.HasLineOfSight(c, blocking) {
			results = append(results, c)
		}
	}

	return results
}

// RectangleGrid returns the set of hexagons that form a rectangle with the specified width and height
func RectangleGrid(width, height int) []Hex {
	var results []Hex

	for q := 0; q < width; q++ {
		qOffset := int(math.Floor(float64(q) / 2.))

		for r := -qOffset; r < height-qOffset; r++ {
			results = append(results, NewHex(q, r))
		}
	}

	return results
}
