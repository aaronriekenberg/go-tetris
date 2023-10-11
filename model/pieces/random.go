package pieces

import (
	"math/rand"

	"github.com/aaronriekenberg/go-tetris/common"
)

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
