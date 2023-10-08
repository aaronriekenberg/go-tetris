package pieces

import "github.com/aaronriekenberg/go-tetris/common"

type TetrisPiece interface {
	CenterCoordinate() common.TetrisModelCoordinate

	CloneWithNewCenterCoordinate(
		newCenterCoordinate common.TetrisModelCoordinate,
	) TetrisPiece

	Coordinates() []common.TetrisModelCoordinate
}

type squarePiece struct {
	centerCoordinate common.TetrisModelCoordinate
}

func (squarePiece squarePiece) CenterCoordinate() common.TetrisModelCoordinate {
	return squarePiece.centerCoordinate
}

func (squarePiece squarePiece) CloneWithNewCenterCoordinate(
	newCenterCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	squarePiece.centerCoordinate = newCenterCoordinate
	return squarePiece
}

func (squarePiece squarePiece) Coordinates() []common.TetrisModelCoordinate {
	return []common.TetrisModelCoordinate{
		{X: squarePiece.centerCoordinate.X, Y: squarePiece.centerCoordinate.Y},
		{X: squarePiece.centerCoordinate.X + 1, Y: squarePiece.centerCoordinate.Y},
		{X: squarePiece.centerCoordinate.X, Y: squarePiece.centerCoordinate.Y + 1},
		{X: squarePiece.centerCoordinate.X + 1, Y: squarePiece.centerCoordinate.Y + 1},
	}
}

func CreateRandomPiece(
	centerCoordinate common.TetrisModelCoordinate,
) TetrisPiece {
	return squarePiece{
		centerCoordinate: centerCoordinate,
	}
}
