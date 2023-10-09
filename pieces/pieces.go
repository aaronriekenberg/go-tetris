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

var pieceConstructors = []func(centerCoordinate common.TetrisModelCoordinate) TetrisPiece{
	newSquarePiece,
	newLinePiece,
}

func CreateRandomPiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return pieceConstructors[rand.Intn(len(pieceConstructors))](centerCoordinate)
}

type squarePiece struct {
	centerCoordinate common.TetrisModelCoordinate

	coordinates []common.TetrisModelCoordinate
}

func newSquarePiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	coordinates := []common.TetrisModelCoordinate{
		centerCoordinate,
		centerCoordinate.AddRows(1),
		centerCoordinate.AddColumns(1),
		centerCoordinate.AddRowsColumns(1, 1),
	}

	return squarePiece{
		centerCoordinate: centerCoordinate,
		coordinates:      coordinates,
	}
}

func (squarePiece squarePiece) Color() tcell.Color {
	return tcell.ColorGreen
}

func (squarePiece squarePiece) CenterCoordinate() common.TetrisModelCoordinate {
	return squarePiece.centerCoordinate
}

func (squarePiece squarePiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newSquarePiece(newCenterCoordinate)
}

func (squarePiece squarePiece) CloneWithNextOrientation() TetrisPiece {
	return newSquarePiece(squarePiece.centerCoordinate)
}

func (squarePiece squarePiece) Coordinates() []common.TetrisModelCoordinate {
	return squarePiece.coordinates
}

type linePiece struct {
	centerCoordinate common.TetrisModelCoordinate

	coordinates []common.TetrisModelCoordinate
}

func newLinePiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	coordinates := []common.TetrisModelCoordinate{
		centerCoordinate,
		centerCoordinate.AddRows(1),
		centerCoordinate.AddRows(2),
		centerCoordinate.AddRows(3),
	}

	return linePiece{
		centerCoordinate: centerCoordinate,
		coordinates:      coordinates,
	}
}

func (linePiece linePiece) Color() tcell.Color {
	return tcell.ColorRed
}

func (linePiece linePiece) CenterCoordinate() common.TetrisModelCoordinate {
	return linePiece.centerCoordinate
}

func (linePiece linePiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newLinePiece(newCenterCoordinate)
}

func (linePiece linePiece) CloneWithNextOrientation() TetrisPiece {
	return newLinePiece(linePiece.centerCoordinate)
}

func (linePiece linePiece) Coordinates() []common.TetrisModelCoordinate {
	return linePiece.coordinates
}
