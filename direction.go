package hexgrid

// Direction defines the sides of a hex
type Direction int

const (
	// DirectionSE is the Southeast side of a hex
	DirectionSE = iota
	// DirectionNE is the Northeast side of a hex
	DirectionNE
	// DirectionN is the North side of a hex
	DirectionN
	// DirectionNW is the Northwest side of a hex
	DirectionNW
	// DirectionSW is the Southwest side of a hex
	DirectionSW
	// DirectionS is the South side of a hex
	DirectionS
)

func (d Direction) ToString() string {
	switch d {
	case DirectionSE:
		return "SE"
	case DirectionNE:
		return "NE"
	case DirectionN:
		return "N"
	case DirectionNW:
		return "NW"
	case DirectionSW:
		return "SW"
	case DirectionS:
		return "S"
	default:
		return "?"
	}
}
