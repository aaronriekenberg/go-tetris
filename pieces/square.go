package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

var createSquareOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(1),
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRowsColumns(1, 1),
		}
	},
}

type squarePiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int
}

func newSquarePiece(
	centerCoordinate common.TetrisModelCoordinate,
	orientation int,
) TetrisPiece {
	return squarePiece{
		centerCoordinate: centerCoordinate,
		orientation:      orientation,
	}
}

func newSquarePieceDefaultOrientation(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newSquarePiece(centerCoordinate, 0)
}

func (squarePiece squarePiece) Color() tcell.Color {
	return tcell.ColorGreen
}

func (squarePiece squarePiece) CenterCoordinate() common.TetrisModelCoordinate {
	return squarePiece.centerCoordinate
}

func (squarePiece squarePiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newSquarePiece(newCenterCoordinate, 0)
}

func (squarePiece squarePiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (squarePiece.orientation + 1) % len(createSquareOrientationFuncs)
	return newSquarePiece(
		squarePiece.centerCoordinate,
		nextOrientation,
	)
}

func (squarePiece squarePiece) Coordinates() []common.TetrisModelCoordinate {
	return createSquareOrientationFuncs[squarePiece.orientation](squarePiece.centerCoordinate)
}
