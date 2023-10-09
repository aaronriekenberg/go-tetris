package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type squarePiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int

	createOrientationFuncs []createOrientationFunc
}

func newSquarePiece(
	centerCoordinate common.TetrisModelCoordinate,
	orientation int,
) TetrisPiece {

	createOrientationFuncs := []createOrientationFunc{
		func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
			return []common.TetrisModelCoordinate{
				centerCoordinate,
				centerCoordinate.AddRows(1),
				centerCoordinate.AddColumns(1),
				centerCoordinate.AddRowsColumns(1, 1),
			}
		},
	}

	return squarePiece{
		centerCoordinate:       centerCoordinate,
		orientation:            orientation,
		createOrientationFuncs: createOrientationFuncs,
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
	return newSquarePiece(newCenterCoordinate, squarePiece.orientation)
}

func (squarePiece squarePiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (squarePiece.orientation + 1) % len(squarePiece.createOrientationFuncs)
	return newSquarePiece(
		squarePiece.centerCoordinate,
		nextOrientation,
	)
}

func (squarePiece squarePiece) Coordinates() []common.TetrisModelCoordinate {
	return squarePiece.createOrientationFuncs[squarePiece.orientation](squarePiece.centerCoordinate)
}
