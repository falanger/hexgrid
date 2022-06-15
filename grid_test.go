package hexgrid

import (
	"testing"
)

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
func TestRectangularGrid(t *testing.T) {
	hexgrid := RectangularGrid(3, 2)
	if len(hexgrid) != 6 {
		t.Error("Expected: 6 got:", len(hexgrid))
	}
}
