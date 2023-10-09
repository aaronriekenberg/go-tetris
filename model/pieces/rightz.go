package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type rigthZPiece struct {
	centerCoordinate common.TetrisModelCoordinate

	orientation int

	createOrientationFuncs []createOrientationFunc
}

var rigthZPieceOrientationFuncs = []createOrientationFunc{
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRows(1),
			centerCoordinate.AddRowsColumns(1, -1),
		}
	},
	func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate {
		return []common.TetrisModelCoordinate{
			centerCoordinate,
			centerCoordinate.AddRows(-1),
			centerCoordinate.AddColumns(1),
			centerCoordinate.AddRowsColumns(1, 1),
		}
	},
}

func newRigthZPiece(
	centerCoordinate common.TetrisModelCoordinate,
	orientation int,
) TetrisPiece {
	return rigthZPiece{
		centerCoordinate:       centerCoordinate,
		orientation:            orientation,
		createOrientationFuncs: rigthZPieceOrientationFuncs,
	}
}

func newRigthZPieceDefaultOrientation(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newRigthZPiece(centerCoordinate, 0)
}

func (rigthZPiece rigthZPiece) Color() tcell.Color {
	return tcell.ColorOrange
}

func (rigthZPiece rigthZPiece) CenterCoordinate() common.TetrisModelCoordinate {
	return rigthZPiece.centerCoordinate
}

func (rigthZPiece rigthZPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newRigthZPiece(newCenterCoordinate, rigthZPiece.orientation)
}

func (rigthZPiece rigthZPiece) CloneWithNextOrientation() TetrisPiece {
	nextOrientation := (rigthZPiece.orientation + 1) % len(rigthZPiece.createOrientationFuncs)
	return newRigthZPiece(
		rigthZPiece.centerCoordinate,
		nextOrientation,
	)
}

func (rigthZPiece rigthZPiece) Coordinates() []common.TetrisModelCoordinate {
	return rigthZPiece.createOrientationFuncs[rigthZPiece.orientation](rigthZPiece.centerCoordinate)
}
