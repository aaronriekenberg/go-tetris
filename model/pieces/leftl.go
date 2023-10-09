package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type leftLPiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int

	createOrientationFuncs []createOrientationFunc
}

var leftLPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRows(2),
			centerCoordinate.AddRowsColumns(2, -1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddColumns(2),
			centerCoordinate.AddRowsColumns(1, 2),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddRows(-2),
			centerCoordinate.AddRowsColumns(-2, 1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(-1),
			centerCoordinate.AddColumns(-2),
			centerCoordinate.AddRowsColumns(-1, -2),
		}
	},
}

func newLeftLPiece(
	centerCoordinate common.TetrisModelCoordinate,
	orientation int,
) TetrisPiece {
	return leftLPiece{
		centerCoordinate:       centerCoordinate,
		orientation:            orientation,
		createOrientationFuncs: leftLPieceOrientationFuncs,
	}
}

func newLeftLPieceDefaultOrientation(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newLeftLPiece(centerCoordinate, 0)
}

func (leftLPiece leftLPiece) Color() tcell.Color {
	return tcell.ColorLightGray
}

func (leftLPiece leftLPiece) CenterCoordinate() common.TetrisModelCoordinate {
	return leftLPiece.centerCoordinate
}

func (leftLPiece leftLPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newLeftLPiece(newCenterCoordinate, leftLPiece.orientation)
}

func (leftLPiece leftLPiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (leftLPiece.orientation + 1) % len(leftLPiece.createOrientationFuncs)
	return newLeftLPiece(
		leftLPiece.centerCoordinate,
		nextOrientation,
	)
}

func (leftLPiece leftLPiece) Coordinates() []common.TetrisModelCoordinate {
	return leftLPiece.createOrientationFuncs[leftLPiece.orientation](leftLPiece.centerCoordinate)
}
