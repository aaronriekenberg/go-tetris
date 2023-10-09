package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type rightLPiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int

	createOrientationFuncs []createOrientationFunc
}

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
	orientation int,
) TetrisPiece {
	return rightLPiece{
		centerCoordinate:       centerCoordinate,
		orientation:            orientation,
		createOrientationFuncs: rightLPieceOrientationFuncs,
	}
}

func newRightLPieceDefaultOrientation(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newRightLPiece(centerCoordinate, 0)
}

func (rightLPiece rightLPiece) Color() tcell.Color {
	return tcell.ColorPink
}

func (rightLPiece rightLPiece) CenterCoordinate() common.TetrisModelCoordinate {
	return rightLPiece.centerCoordinate
}

func (rightLPiece rightLPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newRightLPiece(newCenterCoordinate, rightLPiece.orientation)
}

func (rightLPiece rightLPiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (rightLPiece.orientation + 1) % len(rightLPiece.createOrientationFuncs)
	return newRightLPiece(
		rightLPiece.centerCoordinate,
		nextOrientation,
	)
}

func (rightLPiece rightLPiece) Coordinates() []common.TetrisModelCoordinate {
	return rightLPiece.createOrientationFuncs[rightLPiece.orientation](rightLPiece.centerCoordinate)
}
