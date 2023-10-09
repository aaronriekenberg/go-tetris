package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type leftZPiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int

	createOrientationFuncs []createOrientationFunc
}

func newLeftZPiece(
	centerCoordinate common.TetrisModelCoordinate,
	orientation int,
) TetrisPiece {

	createOrientationFuncs := []createOrientationFunc{
		func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
			return []common.TetrisModelCoordinate{
				centerCoordinate,
				centerCoordinate.AddColumns(-1),
				centerCoordinate.AddRows(1),
				centerCoordinate.AddRowsColumns(1, 1),
			}
		},
		func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
			return []common.TetrisModelCoordinate{
				centerCoordinate,
				centerCoordinate.AddRows(1),
				centerCoordinate.AddColumns(1),
				centerCoordinate.AddRowsColumns(-1, 1),
			}
		},
	}

	return leftZPiece{
		centerCoordinate:       centerCoordinate,
		orientation:            orientation,
		createOrientationFuncs: createOrientationFuncs,
	}
}

func newLeftZPieceDefaultOrientation(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newLeftZPiece(centerCoordinate, 0)
}

func (leftZPiece leftZPiece) Color() tcell.Color {
	return tcell.ColorYellow
}

func (leftZPiece leftZPiece) CenterCoordinate() common.TetrisModelCoordinate {
	return leftZPiece.centerCoordinate
}

func (leftZPiece leftZPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newLeftZPiece(newCenterCoordinate, leftZPiece.orientation)
}

func (leftZPiece leftZPiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (leftZPiece.orientation + 1) % len(leftZPiece.createOrientationFuncs)
	return newLeftZPiece(
		leftZPiece.centerCoordinate,
		nextOrientation,
	)
}

func (leftZPiece leftZPiece) Coordinates() []common.TetrisModelCoordinate {
	return leftZPiece.createOrientationFuncs[leftZPiece.orientation](leftZPiece.centerCoordinate)
}
