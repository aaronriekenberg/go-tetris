package pieces

import (
	"math/rand/v2"

	"github.com/aaronriekenberg/go-tetris/coordinate"
)

type pieceConstructor = func(centerCoordinate coordinate.TetrisModelCoordinate) TetrisPiece

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
	centerCoordinate coordinate.TetrisModelCoordinate,
) TetrisPiece {
	return pieceConstructors[rand.N(len(pieceConstructors))](centerCoordinate)
}
