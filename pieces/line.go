package pieces

import (
	"github.com/gdamore/tcell/v2"

	"github.com/aaronriekenberg/go-tetris/coordinate"
)

var linePieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRows(2),
			centerCoordinate.AddRows(3),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(2),
			centerCoordinate.AddColumns(3),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddRows(-2),
			centerCoordinate.AddRows(-3),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddColumns(-2),
			centerCoordinate.AddColumns(-3),
		}
	},
}

func newLinePiece(
	centerCoordinate coordinate.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorRed,
		centerCoordinate,
		linePieceOrientationFuncs,
	)
}
