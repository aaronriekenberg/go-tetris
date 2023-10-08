package common

const (
	BoardWidth  = 12
	BoardHeight = 16
)

// (0, 0) is topmost and leftmost cell
type TetrisModelCoordinate struct {
	X int
	Y int
}

func (tmc TetrisModelCoordinate) Valid() bool {
	return (tmc.X >= 0) && (tmc.X < BoardWidth) && (tmc.Y >= 0) && (tmc.Y < BoardHeight)
}

func (tmc TetrisModelCoordinate) AddY(delta int) TetrisModelCoordinate {
	tmc.Y += delta
	return tmc
}

func (tmc TetrisModelCoordinate) AddX(delta int) TetrisModelCoordinate {
	tmc.X += delta
	return tmc
}
