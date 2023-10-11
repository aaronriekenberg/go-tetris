package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"

	"github.com/gdamore/tcell/v2"
)

var squarePieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRowsColumns(1, 1),
		}
	},
}

func newSquarePiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorGreen,
		centerCoordinate,
		squarePieceOrientationFuncs,
	)
}
