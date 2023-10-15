package pieces

import (
	"github.com/gdamore/tcell/v2"

	"github.com/aaronriekenberg/go-tetris/coordinate"
)

var rigthZPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRowsColumns(1, -1),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRowsColumns(1, 1),
		}
	},
}

func newRightZPiece(
	centerCoordinate coordinate.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorOrange,
		centerCoordinate,
		rigthZPieceOrientationFuncs,
	)
}
