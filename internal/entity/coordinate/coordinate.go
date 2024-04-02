package coordinate

type Coordinate struct {
	X, Y int
}

func New(x, y int) Coordinate {
	return Coordinate{X: x, Y: y}
}
