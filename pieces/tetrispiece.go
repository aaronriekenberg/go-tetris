package pieces

import (
	"github.com/aaronriekenberg/go-tetris/coord"
	"github.com/gdamore/tcell/v2"
)

type TetrisPiece interface {
	Color() tcell.Color

	CenterCoordinate() coord.TetrisModelCoordinate

	CloneWithNewCenterCoordinate(
		newCenterCoordinate coord.TetrisModelCoordinate,
	) TetrisPiece

	CloneWithNextOrientation() TetrisPiece

	Coordinates() []coord.TetrisModelCoordinate
}

type createOrientationFunc = func(centerCoordinate coord.TetrisModelCoordinate) []coord.TetrisModelCoordinate

type tetrisPiece struct {
	color                  tcell.Color
	centerCoordinate       coord.TetrisModelCoordinate
	orientation            int
	createOrientationFuncs []createOrientationFunc
}

func newTetrisPieceDefaultOrientation(
	color tcell.Color,
	centerCoordinate coord.TetrisModelCoordinate,
	createOrientationFuncs []createOrientationFunc,
) tetrisPiece {
	return tetrisPiece{
		color:                  color,
		centerCoordinate:       centerCoordinate,
		orientation:            0,
		createOrientationFuncs: createOrientationFuncs,
	}
}

func (tetrisPiece tetrisPiece) Color() tcell.Color {
	return tetrisPiece.color
}

func (tetrisPiece tetrisPiece) CenterCoordinate() coord.TetrisModelCoordinate {
	return tetrisPiece.centerCoordinate
}

func (tetrisPiece tetrisPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate coord.TetrisModelCoordinate,
) TetrisPiece {
	tetrisPiece.centerCoordinate = newCenterCoordinate
	return tetrisPiece
}

func (tetrisPiece tetrisPiece) CloneWithNextOrientation() TetrisPiece {
	tetrisPiece.orientation = (tetrisPiece.orientation + 1) % len(tetrisPiece.createOrientationFuncs)
	return tetrisPiece
}

func (tetrisPiece tetrisPiece) Coordinates() []coord.TetrisModelCoordinate {
	return tetrisPiece.createOrientationFuncs[tetrisPiece.orientation](tetrisPiece.centerCoordinate)
}
