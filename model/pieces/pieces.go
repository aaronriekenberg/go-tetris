package pieces

import (
	"math/rand"

	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type TetrisPiece interface {
	Color() tcell.Color

	CenterCoordinate() common.TetrisModelCoordinate

	CloneWithNewCenterCoordinate(
		newCenterCoordinate common.TetrisModelCoordinate,
	) TetrisPiece

	CloneWithNextOrientation() TetrisPiece

	Coordinates() []common.TetrisModelCoordinate
}

type tetrisPiece struct {
	color                  tcell.Color
	centerCoordinate       common.TetrisModelCoordinate
	orientation            int
	createOrientationFuncs []createOrientationFunc
}

func newTetrisPieceDefaultOrientation(
	color tcell.Color,
	centerCoordinate common.TetrisModelCoordinate,
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

func (tetrisPiece tetrisPiece) CenterCoordinate() common.TetrisModelCoordinate {
	return tetrisPiece.centerCoordinate
}

func (tetrisPiece tetrisPiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	tetrisPiece.centerCoordinate = newCenterCoordinate
	return tetrisPiece
}

func (tetrisPiece tetrisPiece) CloneWithNextOrientation() TetrisPiece {
	tetrisPiece.orientation = (tetrisPiece.orientation + 1) % len(tetrisPiece.createOrientationFuncs)
	return tetrisPiece
}

func (tetrisPiece tetrisPiece) Coordinates() []common.TetrisModelCoordinate {
	return tetrisPiece.createOrientationFuncs[tetrisPiece.orientation](tetrisPiece.centerCoordinate)
}

type pieceConstructor = func(centerCoordinate common.TetrisModelCoordinate) TetrisPiece

var pieceConstructors = []pieceConstructor{
	newSquarePiece,
	newLinePiece,
	newTPiece,
	newLeftZPiece,
	newRightZPiece,
	newLeftLPiece,
	newRightLPiece,
}

func CreateRandomPiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return pieceConstructors[rand.Intn(len(pieceConstructors))](centerCoordinate)
}

type createOrientationFunc = func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate
