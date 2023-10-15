package pieces

import (
	"math/rand"

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
	return pieceConstructors[rand.Intn(len(pieceConstructors))](centerCoordinate)
}
