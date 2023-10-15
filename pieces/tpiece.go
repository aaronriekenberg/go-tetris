package pieces

import (
	"github.com/aaronriekenberg/go-tetris/coord"
	"github.com/gdamore/tcell/v2"
)

var tPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddRows(1),
		}
	},
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate.AddRows(-1),
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(1),
		}
	},
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddRows(-1),
		}
	},
	func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate {
		return []coord.TetrisModelCoordinate{
			centerCoordinate.AddRows(-1),
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(-1),
		}
	},
}

func newTPiece(
	centerCoordinate coord.TetrisModelCoordinate,
) TetrisPiece {
	return newTetrisPieceDefaultOrientation(
		tcell.ColorDarkCyan,
		centerCoordinate,
		tPieceOrientationFuncs,
	)
}
