package pieces

import (
	"github.com/gdamore/tcell/v3"

	"github.com/aaronriekenberg/go-tetris/coordinate"
)

var leftLPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRows(2),
			centerCoordinate.AddRowsColumns(2, -1),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(2),
			centerCoordinate.AddRowsColumns(1, 2),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddRows(-2),
			centerCoordinate.AddRowsColumns(-2, 1),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddColumns(-2),
			centerCoordinate.AddRowsColumns(-1, -2),
		}
	},
}

func newLeftLPiece(
	centerCoordinate coordinate.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorLightGray,
		centerCoordinate,
		leftLPieceOrientationFuncs,
	)
}
