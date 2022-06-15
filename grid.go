package hexgrid

import "math"

// RectangularGrid returns the set of hexagons that form a rectangle with the specified width and height
func RectangularGrid(width, height int) []Hex {
	grid := make([]Hex, 0)
	for q := 0; q < width; q++ {
		qOffset := int(math.Floor(float64(q) / 2.))
		for r := -qOffset; r < height-qOffset; r++ {
			grid = append(grid, NewHex(q, r))
		}
	}
	return grid
}

// HexagonalGrid returns a slice of hexagons that form a hexagon with the specified radius.
func HexagonalGrid(radius int) []Hex {
	grid := make([]Hex, 0)
	for q := -radius; q <= radius; q++ {
		r1 := int(math.Max(float64(-radius), float64(-q-radius)))
		r2 := int(math.Min(float64(radius), float64(-q+radius)))
		for r := r1; r <= r2; r++ {
			grid = append(grid, NewHex(q, r))
		}
	}
	return grid
}
