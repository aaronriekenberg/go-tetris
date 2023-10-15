package pieces

import (
	"github.com/aaronriekenberg/go-tetris/coord"
	"github.com/gdamore/tcell/v2"
)

var rightLPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRows(2),
			centerCoordinate.AddRowsColumns(2, 1),
		}
	},
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(2),
			centerCoordinate.AddRowsColumns(-1, 2),
		}
	},
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddRows(-2),
			centerCoordinate.AddRowsColumns(-2, -1),
		}
	},
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddColumns(-2),
			centerCoordinate.AddRowsColumns(1, -2),
		}
	},
}

func newRightLPiece(
	centerCoordinate coord.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorPink,
		centerCoordinate,
		rightLPieceOrientationFuncs,
	)
}
