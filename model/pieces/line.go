package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"

	"github.com/gdamore/tcell/v2"
)

var linePieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRows(2),
			centerCoordinate.AddRows(3),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(2),
			centerCoordinate.AddColumns(3),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddRows(-2),
			centerCoordinate.AddRows(-3),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddColumns(-2),
			centerCoordinate.AddColumns(-3),
		}
	},
}

func newLinePiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorRed,
		centerCoordinate,
		linePieceOrientationFuncs,
	)
}
