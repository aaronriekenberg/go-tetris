package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

var createLineOrientationFuncs = []createOrientationFunc{
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

type linePiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int
}

func newLinePiece(
	centerCoordinate common.TetrisModelCoordinate,
	orientation int,
) TetrisPiece {
	return linePiece{
		centerCoordinate: centerCoordinate,
		orientation:      orientation,
	}
}

func newLinePieceDefaultOrientation(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newLinePiece(centerCoordinate, 0)
}

func (linePiece linePiece) Color() tcell.Color {
	return tcell.ColorRed
}

func (linePiece linePiece) CenterCoordinate() common.TetrisModelCoordinate {
	return linePiece.centerCoordinate
}

func (linePiece linePiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newLinePiece(
		newCenterCoordinate,
		linePiece.orientation,
	)
}

func (linePiece linePiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (linePiece.orientation + 1) % len(createLineOrientationFuncs)
	return newLinePiece(
		linePiece.centerCoordinate,
		nextOrientation,
	)
}

func (linePiece linePiece) Coordinates() []common.TetrisModelCoordinate {
	return createLineOrientationFuncs[linePiece.orientation](linePiece.centerCoordinate)
}
