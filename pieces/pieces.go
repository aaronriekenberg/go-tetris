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

type pieceConstructor = func(centerCoordinate common.TetrisModelCoordinate) TetrisPiece

var pieceConstructors = []pieceConstructor{
	newSquarePieceDefaultOrientation,
	newLinePieceDefaultOrientation,
	newTPieceDefaultOrientation,
	newLeftZPieceDefaultOrientation,
	newRigthZPieceDefaultOrientation,
	newLeftLPieceDefaultOrientation,
}

func CreateRandomPiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return pieceConstructors[rand.Intn(len(pieceConstructors))](centerCoordinate)
}

type createOrientationFunc = func(centerCoordinate common.TetrisModelCoordinate) []common.TetrisModelCoordinate
