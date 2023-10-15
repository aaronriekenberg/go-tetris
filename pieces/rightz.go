package pieces

import (
	"github.com/aaronriekenberg/go-tetris/coord"
	"github.com/gdamore/tcell/v2"
)

var rigthZPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRowsColumns(1, -1),
		}
	},
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRowsColumns(1, 1),
		}
	},
}

func newRightZPiece(
	centerCoordinate coord.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorOrange,
		centerCoordinate,
		rigthZPieceOrientationFuncs,
	)
}
