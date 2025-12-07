package pieces

import (
	"github.com/gdamore/tcell/v3"

	"github.com/aaronriekenberg/go-tetris/coordinate"
)

var squarePieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRowsColumns(1, 1),
		}
	},
}

func newSquarePiece(
	centerCoordinate coordinate.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorGreen,
		centerCoordinate,
		squarePieceOrientationFuncs,
	)
}
