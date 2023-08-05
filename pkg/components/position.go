package components

type Position struct {
	X float64
	Y float64

	CellX int
	CellY int
}

const GridCellWidth = 32
const GridCellHeight = 32

func NewGridPosition(x, y int, cellX, cellY int) Position {
	return Position{
		X:     float64(x) * GridCellWidth,
		Y:     float64(y) * GridCellHeight,
		CellX: cellX,
		CellY: cellY,
	}
}
