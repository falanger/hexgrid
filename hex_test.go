package hexgrid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexAdd(t *testing.T) {
	var testCases = []struct {
		hexA     Hex
		hexB     Hex
		expected Hex
	}{
		{NewHex(1, -3), NewHex(3, -7), NewHex(4, -10)},
	}

	for _, tt := range testCases {
		actual := tt.hexA.Add(tt.hexB)
		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

func TestHexSubtract(t *testing.T) {
	var testCases = []struct {
		hexA     Hex
		hexB     Hex
		expected Hex
	}{
		{NewHex(1, -3), NewHex(3, -7), NewHex(-2, 4)},
	}

	for _, tt := range testCases {
		actual := tt.hexA.Subtract(tt.hexB)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

func TestHexScale(t *testing.T) {
	var testCases = []struct {
		hexA     Hex
		factor   int
		expected Hex
	}{
		{NewHex(1, -3), 2, NewHex(2, -6)},
		{NewHex(-2, 3), 2, NewHex(-4, 6)},
	}

	for _, tt := range testCases {
		actual := tt.hexA.Scale(tt.factor)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}

}

//           _ _
//         /     \
//    _ _ /(0,-2) \ _ _
//  /     \       /     \
// /(-1,-1)\ _ _ /(1,-2) \
// \       /     \       /
//  \ _ _ /(0,-1) \ _ _ /
//  /     \       /     \
// /(-1,0) \ _ _ /(1,-1) \
// \       /     \       /
//  \ _ _ / (0,0) \ _ _ /
//        \       /
//         \ _ _ /
// Tests that the neighbors of a certain hexagon are properly computed for all directions
func TestHexNeighbor(t *testing.T) {

	var testCases = []struct {
		origin    Hex
		direction Direction
		expected  Hex
	}{

		{NewHex(0, -1), DirectionSE, NewHex(1, -1)},
		{NewHex(0, -1), DirectionNE, NewHex(1, -2)},
		{NewHex(0, -1), DirectionN, NewHex(0, -2)},
		{NewHex(0, -1), DirectionNW, NewHex(-1, -1)},
		{NewHex(0, -1), DirectionSW, NewHex(-1, 0)},
		{NewHex(0, -1), DirectionS, NewHex(0, 0)},
	}

	for _, tt := range testCases {
		actual := tt.origin.Neighbor(tt.direction)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

// DISTANCE TESTS

//           _ _
//         /     \
//    _ _ /(0,-2) \ _ _
//  /     \       /     \
// /(-1,-1)\ _ _ /(1,-2) \
// \       /     \       /
//  \ _ _ /(0,-1) \ _ _ /
//  /     \       /     \
// /(-1,0) \ _ _ /(1,-1) \
// \       /     \       /
//  \ _ _ / (0,0) \ _ _ /
//  /     \       /     \
// /(-1,1) \ _ _ / (1,0) \
// \       /     \       /
//  \ _ _ / (0,1) \ _ _ /
//        \       /
//         \ _ _ /

func TestHexDistance(t *testing.T) {
	var testCases = []struct {
		origin      Hex
		destination Hex
		expected    int
	}{
		{NewHex(-1, -1), NewHex(1, -1), 2},
		{NewHex(-1, -1), NewHex(0, 0), 2},
		{NewHex(0, -1), NewHex(0, -2), 1},
		{NewHex(-1, -1), NewHex(0, 1), 3},
		{NewHex(1, 0), NewHex(-1, -1), 3},
	}

	for _, tt := range testCases {
		actual := tt.origin.DistanceTo(tt.destination)

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

//          _____         _____         _____
//         /     \       /     \       /     \
//   _____/ -2,-2 \_____/  0,-3 \_____/  2,-4 \_____
//  /     \       /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \_____/  3,-4 \
// \       /     \       /     \       /     \       /
//  \_____/ -2,-1 \_____/  0,-2 \_____/  2,-3 \_____/
//  /     \       /     \       /     \       /     \
// / -3,0  \_____/ -1,-1 \_____/  1,-2 \_____/  3,-3 \
// \       /     \       /     \       /     \       /
//  \_____/ -2,0  \_____/  0,-1 \_____/  2,-2 \_____/
//  /     \       /     \       /     \       /     \
// / -3,1  \_____/ -1,0  \_____/  1,-1 \_____/  3,-2 \
// \       /     \       /     \       /     \       /
//  \_____/       \_____/       \_____/       \_____/
func TestHexLineDraw(t *testing.T) {
	var testCases = []struct {
		origin      Hex
		destination Hex
		expected    string // the expected path serialized to string
	}{
		{NewHex(-3, -1), NewHex(3, -3), "[(-3,-1) (-2,-1) (-1,-2) (0,-2) (1,-2) (2,-3) (3,-3)]"},
		{NewHex(-2, 0), NewHex(2, -2), "[(-2,0) (-1,0) (0,-1) (1,-1) (2,-2)]"},
		{NewHex(1, -1), NewHex(1, -3), "[(1,-1) (1,-2) (1,-3)]"},
	}

	for _, tt := range testCases {
		actual := fmt.Sprint(tt.origin.LineDraw(tt.destination))

		if actual != tt.expected {
			t.Error("Expected:", tt.expected, "got:", actual)
		}
	}
}

// Tests that the range includes the correct number of hexagons with a certain radius from the center
//                 _____
//                /     \
//          _____/ -1,-2 \_____
//         /     \       /     \
//   _____/ -2,-1 \_____/  0,-2 \_____
//  /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \
// \       /     \       /     \       /
//  \_____/ -2,-2 \_____/  0,-3 \_____/
//  /     \       /     \       /     \
// / -3,-1 \_____/ -1,-2 \_____/  1,-3 \
// \       /     \ CENTR /     \       /
//  \_____/ -2,-1 \_____/  0,-2 \_____/
//  /     \       /     \       /     \
// / -3,0  \_____/ -1,-1 \_____/  1,-2 \
// \       /     \       /     \       /
//  \_____/ -2,0  \_____/  0,-1 \_____/
//        \       /     \       /
//         \_____/ -1,0  \_____/
//               \       /
//                \_____/
func TestHexRange(t *testing.T) {
	var testCases = []struct {
		radius                   int
		expectedNumberOfHexagons int
	}{
		{0, 1},
		{1, 7},
		{2, 19},
	}

	center := NewHex(1, -2)

	for _, tt := range testCases {
		actual := center.Range(tt.radius)

		if len(actual) != tt.expectedNumberOfHexagons {
			t.Error("Expected:", tt.expectedNumberOfHexagons, "got:", len(actual))
		}
	}
}

//    _ _           _ _
//  /     \       /     \
// /( 0,0) \ _ _ /(2,-1) \
// \       /     \       /
//  \ _ _ / (1,0) \ _ _ /
//  /     \       /     \
// / (0,1) \ _ _ / (2,0) \
// \       /     \       /
//  \ _ _ / (1,1) \ _ _ /
//        \       /
//         \ _ _ /
func TestHexRectangle(t *testing.T) {
	hexgrid := RectangleGrid(3, 2)
	expectHexes := 6
	assert.Len(t, hexgrid, expectHexes)
}

//    _ _           _ _           _ _
//  /     \       /     \       /     \
// /  0 0  \ _ _ /  2-1  \ _ _ /  4-2  \ _ _
// \       /     \   X   /     \   X   /     \
//  \ _ _ /  1 0  \ _ _ /  3-1  \ _ _ /  5-2  \
//  /     \       /# # #\   X   /     \   X   /
// /  0 1  \ _ _ /# 2 0 #\ _ _ /  4-1  \ _ _ /
// \       /     \#     #/# # #\   X   /     \
//  \ _ _ /  1 1  \#_#_#/# 3 0 #\ _ _ /  5-1  \
//  /     \  |P|  /     \#  X  #/     \   X   /
// /  0 2  \ _ _ /  2 1  \#_#_#/  4 0  \ _ _ /
// \       /     \       /     \   X   /     \
//  \ _ _ /  1 2  \ _ _ /  3 1  \ _ _ /  5 0  \
//  /     \       /     \       /     \       /
// /  0 3  \ _ _ /  2 2  \ _ _ /  4 1  \ _ _ /
// \       /     \       /     \       /     \
//  \ _ _ /  1 3  \ _ _ /  3 2  \ _ _ /  5 1  \
//        \       /     \       /     \       /
//         \ _ _ /       \ _ _ /       \ _ _ /
//
// The FOV measured from the central hex at 1,1, assuming blocking hexagons at 2,0 and 3,0.
// The hexagons marked with an X are non-visible. The remaining 16 are visible.
func TestHexFieldOfView(t *testing.T) {

	universe := []Hex{
		NewHex(0, 0),
		NewHex(0, 1),
		NewHex(0, 2),
		NewHex(0, 3),
		NewHex(1, 0),
		NewHex(1, 1),
		NewHex(1, 2),
		NewHex(1, 3),
		NewHex(2, -1),
		NewHex(2, 0),
		NewHex(2, 1),
		NewHex(2, 2),
		NewHex(3, -1),
		NewHex(3, 0),
		NewHex(3, 1),
		NewHex(3, 2),
		NewHex(4, -2),
		NewHex(4, -1),
		NewHex(4, 0),
		NewHex(4, 1),
		NewHex(5, -2),
		NewHex(5, -1),
		NewHex(5, 0),
		NewHex(5, 1),
	}

	losBlockers := []Hex{NewHex(2, 0), NewHex(3, 0)}

	point := NewHex(1, 1)
	actual := point.FieldOfView(universe, losBlockers)

	if len(actual) != 16 {
		t.Error("Expected: 16 got:", len(actual))
	}
}

////////////////
// Benchmarks //
////////////////

func BenchmarkHexDistance(b *testing.B) {
	var testCases = []struct {
		destination Hex
	}{
		{NewHex(0, 0)},
		{NewHex(100, 100)},
		{NewHex(10000, 10000)},
	}

	for _, bm := range testCases {

		origin := NewHex(0, 0)

		b.Run(fmt.Sprint(origin, ":", bm.destination), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				origin.DistanceTo(bm.destination)
			}
		})
	}
}

func BenchmarkHexLineDraw(b *testing.B) {
	var testCases = []struct {
		destination Hex
	}{
		{NewHex(0, 0)},
		{NewHex(100, 100)},
		{NewHex(10000, 10000)},
	}

	for _, bm := range testCases {
		origin := NewHex(0, 0)

		b.Run(fmt.Sprint(origin, ":", bm.destination), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				origin.LineDraw(bm.destination)
			}
		})
	}
}

func BenchmarkHexRange(b *testing.B) {
	var testCases = []struct {
		radius int
	}{
		{0},
		{10},
		{100},
	}

	for _, bm := range testCases {

		origin := NewHex(0, 0)

		b.Run(fmt.Sprint(origin, ":", bm.radius), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				center := NewHex(1, -2)
				center.Range(bm.radius)
			}
		})
	}
}

func BenchmarkHexHasLineOfSight(b *testing.B) {
	center := NewHex(1, 1)
	for i := 0; i < b.N; i++ {
		center.HasLineOfSight(NewHex(4, -1), []Hex{NewHex(2, 0), NewHex(3, 0)})
	}
}
