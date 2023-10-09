package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type tPiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int

	createOrientationFuncs []createOrientationFunc
}

var tPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddRows(1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate.AddRows(-1),
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddRows(-1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate.AddRows(-1),
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(-1),
		}
	},
}

func newTPiece(
	centerCoordinate common.TetrisModelCoordinate,
	orientation int,
) TetrisPiece {

	return tPiece{
		centerCoordinate:       centerCoordinate,
		orientation:            orientation,
		createOrientationFuncs: tPieceOrientationFuncs,
	}
}

func newTPieceDefaultOrientation(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newTPiece(centerCoordinate, 0)
}

func (tPiece tPiece) Color() tcell.Color {
	return tcell.ColorDarkCyan
}

func (tPiece tPiece) CenterCoordinate() common.TetrisModelCoordinate {
	return tPiece.centerCoordinate
}

func (tPiece tPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newTPiece(
		newCenterCoordinate,
		tPiece.orientation,
	)
}

func (tPiece tPiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (tPiece.orientation + 1) % len(tPiece.createOrientationFuncs)
	return newTPiece(
		tPiece.centerCoordinate,
		nextOrientation,
	)
}

func (tPiece tPiece) Coordinates() []common.TetrisModelCoordinate {
	return tPiece.createOrientationFuncs[tPiece.orientation](tPiece.centerCoordinate)
}
