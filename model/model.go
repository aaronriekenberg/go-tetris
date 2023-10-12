package model

import (
	"slices"

	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/aaronriekenberg/go-tetris/pieces"

	"github.com/gdamore/tcell/v2"
)

type TetrisModelCell interface {
	Occupied() bool
	Color() tcell.Color
}

type tetrisModelCell struct {
	occupied bool
	color    tcell.Color
}

func (tmc tetrisModelCell) Occupied() bool {
	return tmc.occupied
}

func (tmc tetrisModelCell) Color() tcell.Color {
	return tmc.color
}

type DrawableInfoModel interface {
	DrawableCells() [][]TetrisModelCell
	Lines() int
	GameOver() bool
}

type TetrisModel interface {
	DrawableInfoModel
	MoveCurrentPieceDown()
	MoveCurrentPieceLeft()
	MoveCurrentPieceRight()
	RotateCurrentPiece()
	DropCurrentPiece()
	PeriodicUpdate()
}

type tetrisModel struct {
	drawableCellsCache [][]TetrisModelCell
	currentPiece       pieces.TetrisPiece
	stackCells         [][]tetrisModelCell
	lines              int
	gameOver           bool
}

func NewTetrisModel() TetrisModel {
	tetrisModel := &tetrisModel{}

	tetrisModel.initializeStackCells()

	return tetrisModel
}

func (tetrisModel *tetrisModel) initializeStackCells() {
	stackCells := make([][]tetrisModelCell, common.BoardRows)

	for row := 0; row < common.BoardRows; row += 1 {
		stackCells[row] = make([]tetrisModelCell, common.BoardColumns)
	}

	tetrisModel.stackCells = stackCells
}

func (tetrisModel *tetrisModel) DrawableCells() [][]TetrisModelCell {
	if tetrisModel.drawableCellsCache != nil {
		return tetrisModel.drawableCellsCache
	}

	drawableCellsCache := make([][]TetrisModelCell, common.BoardRows)

	for row := 0; row < common.BoardRows; row += 1 {
		drawableCellsCache[row] = make([]TetrisModelCell, common.BoardColumns)
		for column := 0; column < common.BoardColumns; column += 1 {
			drawableCellsCache[row][column] = tetrisModel.stackCells[row][column]
		}
	}

	if tetrisModel.currentPiece != nil {
		for _, coordinates := range tetrisModel.currentPiece.Coordinates() {
			drawableCell := tetrisModelCell{
				occupied: true,
				color:    tetrisModel.currentPiece.Color(),
			}
			drawableCellsCache[coordinates.Row()][coordinates.Column()] = drawableCell
		}
	}

	tetrisModel.drawableCellsCache = drawableCellsCache

	return tetrisModel.drawableCellsCache
}

func (tetrisModel *tetrisModel) Lines() int {
	return tetrisModel.lines
}

func (tetrisModel *tetrisModel) GameOver() bool {
	return tetrisModel.gameOver
}

func (tetrisModel *tetrisModel) invalidateDrawableCellsCache() {
	tetrisModel.drawableCellsCache = nil
}

func (tetrisModel *tetrisModel) isPieceLocationValid(
	tetrisPiece pieces.TetrisPiece,
) bool {
	for _, coordinate := range tetrisPiece.Coordinates() {
		if !coordinate.Valid() {
			return false
		} else if tetrisModel.stackCells[coordinate.Row()][coordinate.Column()].occupied {
			return false
		}
	}
	return true
}

func (tetrisModel *tetrisModel) addNewPiece() {
	centerCoordinate := common.NewTetrisModelCoordinate(
		0,
		(common.BoardColumns/2)-1,
	)

	newPiece := pieces.CreateRandomPiece(centerCoordinate)

	if !tetrisModel.isPieceLocationValid(newPiece) {
		tetrisModel.gameOver = true
		return
	}

	tetrisModel.currentPiece = newPiece
}

func (tetrisModel *tetrisModel) MoveCurrentPieceDown() {
	if tetrisModel.gameOver {
		return
	}

	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddRows(1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if !tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.addCurrentPieceToStack()
		} else {
			tetrisModel.currentPiece = updatedPiece
		}

		tetrisModel.invalidateDrawableCellsCache()
	}
}

func (tetrisModel *tetrisModel) MoveCurrentPieceLeft() {
	if tetrisModel.gameOver {
		return
	}

	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddColumns(-1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.currentPiece = updatedPiece

			tetrisModel.invalidateDrawableCellsCache()
		}
	}
}

func (tetrisModel *tetrisModel) MoveCurrentPieceRight() {
	if tetrisModel.gameOver {
		return
	}

	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddColumns(1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.currentPiece = updatedPiece

			tetrisModel.invalidateDrawableCellsCache()
		}
	}
}

func (tetrisModel *tetrisModel) RotateCurrentPiece() {
	if tetrisModel.gameOver {
		return
	}

	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		updatedPiece := currentPiece.CloneWithNextOrientation()

		if tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.currentPiece = updatedPiece

			tetrisModel.invalidateDrawableCellsCache()
		}
	}
}

func (tetrisModel *tetrisModel) DropCurrentPiece() {
	if tetrisModel.gameOver {
		return
	}

	for tetrisModel.currentPiece != nil {
		tetrisModel.MoveCurrentPieceDown()
	}
}

func (tetrisModel *tetrisModel) addCurrentPieceToStack() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		for _, coordinate := range currentPiece.Coordinates() {
			tetrisModel.stackCells[coordinate.Row()][coordinate.Column()].occupied = true
			tetrisModel.stackCells[coordinate.Row()][coordinate.Column()].color = currentPiece.Color()
		}
	}
	tetrisModel.currentPiece = nil

	tetrisModel.handleFilledStackRows()
}

func (tetrisModel *tetrisModel) handleFilledStackRows() {
	row := common.BoardRows - 1

	for row >= 0 {
		rowIsFull := true
		for _, cell := range tetrisModel.stackCells[row] {
			if !cell.occupied {
				rowIsFull = false
				break
			}
		}
		if rowIsFull {
			tetrisModel.stackCells = slices.Delete(tetrisModel.stackCells, row, row+1)
			tetrisModel.stackCells = slices.Insert(tetrisModel.stackCells, 0, make([]tetrisModelCell, common.BoardColumns))
			tetrisModel.lines += 1
		} else {
			row -= 1
		}
	}
}

func (tetrisModel *tetrisModel) PeriodicUpdate() {
	if tetrisModel.gameOver {
		return
	}

	if tetrisModel.currentPiece == nil {
		tetrisModel.addNewPiece()
	} else {
		tetrisModel.MoveCurrentPieceDown()
	}

	tetrisModel.invalidateDrawableCellsCache()
}
