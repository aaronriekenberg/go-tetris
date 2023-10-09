package common

const (
	BoardRows    = 16
	BoardColumns = 12
)

// (0, 0) is topmost and leftmost cell
type TetrisModelCoordinate struct {
	row    int
	column int
}

func NewTetrisModelCoordinate(
	row, column int,
) TetrisModelCoordinate {
	return TetrisModelCoordinate{
		row:    row,
		column: column,
	}
}

func (tmc TetrisModelCoordinate) Valid() bool {
	return (tmc.row >= 0) && (tmc.row < BoardRows) && (tmc.column >= 0) && (tmc.column < BoardColumns)
}

func (tmc TetrisModelCoordinate) AddRows(delta int) TetrisModelCoordinate {
	tmc.row += delta
	return tmc
}

func (tmc TetrisModelCoordinate) AddColumns(delta int) TetrisModelCoordinate {
	tmc.column += delta
	return tmc
}

func (tmc TetrisModelCoordinate) AddRowsColumns(
	rowsDelta, columnsDelta int,
) TetrisModelCoordinate {
	tmc.row += rowsDelta
	tmc.column += columnsDelta
	return tmc
}

func (tmc TetrisModelCoordinate) Row() int {
	return tmc.row
}

func (tmc TetrisModelCoordinate) Column() int {
	return tmc.column
}
