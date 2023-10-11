package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"

	"github.com/gdamore/tcell/v2"
)

var leftZPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRowsColumns(1, 1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRowsColumns(-1, 1),
		}
	},
}

func newLeftZPiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorYellow,
		centerCoordinate,
		leftZPieceOrientationFuncs,
	)
}
