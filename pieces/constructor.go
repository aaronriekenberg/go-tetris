package pieces

import (
	"math/rand"

	"github.com/aaronriekenberg/go-tetris/coord"
)

type pieceConstructor = func(centerCoordinate coord.TetrisModelCoordinate) TetrisPiece

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
	centerCoordinate coord.TetrisModelCoordinate,
) TetrisPiece {
	return pieceConstructors[rand.Intn(len(pieceConstructors))](centerCoordinate)
}
