package pieces

import (
	"github.com/gdamore/tcell/v3"

	"github.com/aaronriekenberg/go-tetris/coordinate"
)

var tPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddRows(1),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate.AddRows(-1),
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(1),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddRows(-1),
		}
	},
	func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate {
		return []coordinate.TetrisModelCoordinate{
			centerCoordinate.AddRows(-1),
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(-1),
		}
	},
}

func newTPiece(
	centerCoordinate coordinate.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorDarkCyan,
		centerCoordinate,
		tPieceOrientationFuncs,
	)
}
