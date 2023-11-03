package model

import (
	"slices"

	"github.com/aaronriekenberg/go-tetris/coordinate"
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
	Restart()
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
	return newTetrisModel()
}

func newTetrisModel() *tetrisModel {
	tetrisModel := &tetrisModel{
		stackCells: createStackCells(),
	}

	return tetrisModel
}

func createStackCells() (stackCells [][]tetrisModelCell) {
	stackCells = make([][]tetrisModelCell, coordinate.BoardModelRows)

	for row := 0; row < coordinate.BoardModelRows; row += 1 {
		stackCells[row] = make([]tetrisModelCell, coordinate.BoardModelColumns)
	}

	return
}

func (tm *tetrisModel) DrawableCells() [][]TetrisModelCell {
	if tm.drawableCellsCache != nil {
		return tm.drawableCellsCache
	}

	tm.drawableCellsCache = make([][]TetrisModelCell, coordinate.BoardModelRows)

	for row := 0; row < coordinate.BoardModelRows; row += 1 {
		tm.drawableCellsCache[row] = make([]TetrisModelCell, coordinate.BoardModelColumns)
		for column := 0; column < coordinate.BoardModelColumns; column += 1 {
			tm.drawableCellsCache[row][column] = tm.stackCells[row][column]
		}
	}

	if tm.currentPiece != nil {
		for _, coordinates := range tm.currentPiece.Coordinates() {
			drawableCell := tetrisModelCell{
				occupied: true,
				color:    tm.currentPiece.Color(),
			}
			tm.drawableCellsCache[coordinates.Row()][coordinates.Column()] = drawableCell
		}
	}

	return tm.drawableCellsCache
}

func (tm *tetrisModel) invalidateDrawableCellsCache() {
	tm.drawableCellsCache = nil
}

func (tm *tetrisModel) Lines() int {
	return tm.lines
}

func (tm *tetrisModel) GameOver() bool {
	return tm.gameOver
}

func (tm *tetrisModel) isPieceLocationValid(
	tetrisPiece pieces.TetrisPiece,
) bool {
	for _, coordinate := range tetrisPiece.Coordinates() {
		if !coordinate.Valid() {
			return false
		} else if tm.stackCells[coordinate.Row()][coordinate.Column()].occupied {
			return false
		}
	}
	return true
}

func (tm *tetrisModel) addNewPiece() {
	centerCoordinate := coordinate.NewTetrisModelCoordinate(
		0,
		(coordinate.BoardModelColumns/2)-1,
	)

	newPiece := pieces.CreateRandomPiece(centerCoordinate)

	if !tm.isPieceLocationValid(newPiece) {
		tm.gameOver = true
		return
	}

	tm.currentPiece = newPiece
}

func (tm *tetrisModel) Restart() {
	*tm = *newTetrisModel()
}

func (tm *tetrisModel) MoveCurrentPieceDown() {
	if tm.gameOver {
		return
	}

	currentPiece := tm.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddRows(1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if !tm.isPieceLocationValid(updatedPiece) {
			tm.addCurrentPieceToStack()
		} else {
			tm.currentPiece = updatedPiece
		}

		tm.invalidateDrawableCellsCache()
	}
}

func (tm *tetrisModel) MoveCurrentPieceLeft() {
	if tm.gameOver {
		return
	}

	currentPiece := tm.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddColumns(-1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if tm.isPieceLocationValid(updatedPiece) {
			tm.currentPiece = updatedPiece

			tm.invalidateDrawableCellsCache()
		}
	}
}

func (tm *tetrisModel) MoveCurrentPieceRight() {
	if tm.gameOver {
		return
	}

	currentPiece := tm.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddColumns(1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if tm.isPieceLocationValid(updatedPiece) {
			tm.currentPiece = updatedPiece

			tm.invalidateDrawableCellsCache()
		}
	}
}

func (tm *tetrisModel) RotateCurrentPiece() {
	if tm.gameOver {
		return
	}

	currentPiece := tm.currentPiece
	if currentPiece != nil {
		updatedPiece := currentPiece.CloneWithNextOrientation()

		if tm.isPieceLocationValid(updatedPiece) {
			tm.currentPiece = updatedPiece

			tm.invalidateDrawableCellsCache()
		}
	}
}

func (tm *tetrisModel) DropCurrentPiece() {
	if tm.gameOver {
		return
	}

	for tm.currentPiece != nil {
		tm.MoveCurrentPieceDown()
	}
}

func (tm *tetrisModel) addCurrentPieceToStack() {
	currentPiece := tm.currentPiece
	if currentPiece != nil {
		for _, coordinate := range currentPiece.Coordinates() {
			stackCell := &tm.stackCells[coordinate.Row()][coordinate.Column()]
			stackCell.occupied = true
			stackCell.color = currentPiece.Color()
		}
	}
	tm.currentPiece = nil

	tm.handleFilledStackRows()
}

func (tm *tetrisModel) handleFilledStackRows() {
	rowIsFull := func(row int) bool {
		for _, cell := range tm.stackCells[row] {
			if !cell.occupied {
				return false
			}
		}
		return true
	}

	modifiedStackCells := false
	for row := coordinate.BoardModelRows - 1; row >= 0; {
		if rowIsFull(row) {
			tm.stackCells = slices.Delete(tm.stackCells, row, row+1)
			tm.stackCells = slices.Insert(tm.stackCells, 0, make([]tetrisModelCell, coordinate.BoardModelColumns))
			modifiedStackCells = true
			tm.lines += 1
		} else {
			row -= 1
		}
	}

	if modifiedStackCells {
		tm.stackCells = slices.Clip(tm.stackCells)
	}
}

func (tm *tetrisModel) PeriodicUpdate() {
	if tm.gameOver {
		return
	}

	if tm.currentPiece == nil {
		tm.addNewPiece()
	} else {
		tm.MoveCurrentPieceDown()
	}

	tm.invalidateDrawableCellsCache()
}
