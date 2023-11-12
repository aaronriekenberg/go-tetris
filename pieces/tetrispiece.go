package pieces

import (
	"github.com/gdamore/tcell/v2"

	"github.com/aaronriekenberg/go-tetris/coordinate"
)

type TetrisPiece interface {
	Color() tcell.Color

	CenterCoordinate() coordinate.TetrisModelCoordinate

	CloneWithNewCenterCoordinate(
		newCenterCoordinate coordinate.TetrisModelCoordinate,
	) TetrisPiece

	CloneWithNextOrientation() TetrisPiece

	Coordinates() []coordinate.TetrisModelCoordinate
}

type createOrientationFunc = func(centerCoordinate coordinate.TetrisModelCoordinate) []coordinate.TetrisModelCoordinate

// Immutable after creation.
type tetrisPiece struct {
	color                  tcell.Color
	centerCoordinate       coordinate.TetrisModelCoordinate
	orientation            int
	createOrientationFuncs []createOrientationFunc
}

func newTetrisPieceDefaultOrientation(
	color tcell.Color,
	centerCoordinate coordinate.TetrisModelCoordinate,
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

func (tetrisPiece tetrisPiece) CenterCoordinate() coordinate.TetrisModelCoordinate {
	return tetrisPiece.centerCoordinate
}

func (tetrisPiece tetrisPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate coordinate.TetrisModelCoordinate,
) TetrisPiece {
	tetrisPiece.centerCoordinate = newCenterCoordinate
	return tetrisPiece
}

func (tetrisPiece tetrisPiece) CloneWithNextOrientation() TetrisPiece {
	tetrisPiece.orientation = (tetrisPiece.orientation + 1) % len(tetrisPiece.createOrientationFuncs)
	return tetrisPiece
}

func (tetrisPiece tetrisPiece) Coordinates() []coordinate.TetrisModelCoordinate {
	return tetrisPiece.createOrientationFuncs[tetrisPiece.orientation](tetrisPiece.centerCoordinate)
}
