package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"

	"github.com/gdamore/tcell/v2"
)

var rightLPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRows(2),
			centerCoordinate.AddRowsColumns(2, 1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(2),
			centerCoordinate.AddRowsColumns(-1, 2),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddRows(-2),
			centerCoordinate.AddRowsColumns(-2, -1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddColumns(-2),
			centerCoordinate.AddRowsColumns(1, -2),
		}
	},
}

func newRightLPiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorPink,
		centerCoordinate,
		rightLPieceOrientationFuncs,
	)
}
