package common

const (
	BoardWidth  = 12
	BoardHeight = 16
)

// (0, 0) is topmost and leftmost cell
type TetrisModelCoordinate struct {
	x int
	y int
}

func NewTetrisModelCoordinate(
	x, y int,
) TetrisModelCoordinate {
	return TetrisModelCoordinate{
		x: x,
		y: y,
	}
}

func (tmc TetrisModelCoordinate) Valid() bool {
	return (tmc.x >= 0) && (tmc.x < BoardWidth) && (tmc.y >= 0) && (tmc.y < BoardHeight)
}

func (tmc TetrisModelCoordinate) AddY(delta int) TetrisModelCoordinate {
	tmc.y += delta
	return tmc
}

func (tmc TetrisModelCoordinate) AddX(delta int) TetrisModelCoordinate {
	tmc.x += delta
	return tmc
}

func (tmc TetrisModelCoordinate) X() int {
	return tmc.x
}

func (tmc TetrisModelCoordinate) Y() int {
	return tmc.y
}
