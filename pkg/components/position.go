package components

type Position struct {
	X float64
	Y float64

	CellX int
	CellY int
}

const GridWidth = 20
const GridHeight = 15

func NewGridPosition(x, y int, cellX, cellY int) Position {
	return Position{
		X:     float64(x) / GridWidth,
		Y:     float64(y) / GridHeight,
		CellX: cellX,
		CellY: cellY,
	}
}
