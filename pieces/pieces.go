package pieces

import (
	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/gdamore/tcell/v2"
)

type TetrisPiece interface {
	Color() tcell.Color

	CenterCoordinate() common.TetrisModelCoordinate

	CloneWithNewCenterCoordinate(
		newCenterCoordinate common.TetrisModelCoordinate,
	) TetrisPiece

	Coordinates() []common.TetrisModelCoordinate
}

type squarePiece struct {
	centerCoordinate common.TetrisModelCoordinate

	coordinates []common.TetrisModelCoordinate
}

func newSquarePiece(
	centerCoordinate common.TetrisModelCoordinate,
) squarePiece {
	coordinates := []common.TetrisModelCoordinate{
		common.NewTetrisModelCoordinate(centerCoordinate.Row(), centerCoordinate.Column()),
		common.NewTetrisModelCoordinate(centerCoordinate.Row()+1, centerCoordinate.Column()),
		common.NewTetrisModelCoordinate(centerCoordinate.Row(), centerCoordinate.Column()+1),
		common.NewTetrisModelCoordinate(centerCoordinate.Row()+1, centerCoordinate.Column()+1),
	}

	return squarePiece{
		centerCoordinate: centerCoordinate,
		coordinates:      coordinates,
	}
}

func (squarePiece squarePiece) Color() tcell.Color {
	return tcell.ColorBlue
}

func (squarePiece squarePiece) CenterCoordinate() common.TetrisModelCoordinate {
	return squarePiece.centerCoordinate
}

func (squarePiece squarePiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newSquarePiece(newCenterCoordinate)
}

func (squarePiece squarePiece) Coordinates() []common.TetrisModelCoordinate {
	return squarePiece.coordinates
}

func CreateRandomPiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return newSquarePiece(centerCoordinate)
}
