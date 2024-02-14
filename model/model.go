package model

import (
	"slices"
	"time"

	"github.com/aaronriekenberg/go-tetris/coordinate"
	"github.com/aaronriekenberg/go-tetris/pieces"

	"github.com/gdamore/tcell/v2"
)

type tetrisModelCell struct {
	occupied bool
	color    tcell.Color
}

type DrawableCellsMap = map[coordinate.TetrisModelCoordinate]tcell.Color

type DrawableInfoModel interface {
	DrawableCells() DrawableCellsMap
	Lines() int
	GameOver() bool
	UpdateDuration() time.Duration
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
	drawableCellsCache drawableCellsCache
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

	for row := range stackCells {
		stackCells[row] = make([]tetrisModelCell, coordinate.BoardModelColumns)
	}

	return
}

func (tm *tetrisModel) DrawableCells() DrawableCellsMap {
	return tm.drawableCellsCache.drawableCells(tm)
}

func (tm *tetrisModel) invalidateDrawableCellsCache() {
	tm.drawableCellsCache.invalidate()
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
	if currentPiece == nil {
		return
	}

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

func (tm *tetrisModel) MoveCurrentPieceLeft() {
	if tm.gameOver {
		return
	}

	currentPiece := tm.currentPiece
	if currentPiece == nil {
		return
	}

	centerCoordinate := currentPiece.CenterCoordinate()

	newCenterCoordinate := centerCoordinate.AddColumns(-1)

	updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

	if tm.isPieceLocationValid(updatedPiece) {
		tm.currentPiece = updatedPiece

		tm.invalidateDrawableCellsCache()
	}
}

func (tm *tetrisModel) MoveCurrentPieceRight() {
	if tm.gameOver {
		return
	}

	currentPiece := tm.currentPiece
	if currentPiece == nil {
		return
	}

	centerCoordinate := currentPiece.CenterCoordinate()

	newCenterCoordinate := centerCoordinate.AddColumns(1)

	updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

	if tm.isPieceLocationValid(updatedPiece) {
		tm.currentPiece = updatedPiece

		tm.invalidateDrawableCellsCache()
	}
}

func (tm *tetrisModel) RotateCurrentPiece() {
	if tm.gameOver {
		return
	}

	currentPiece := tm.currentPiece
	if currentPiece == nil {
		return
	}

	updatedPiece := currentPiece.CloneWithNextOrientation()

	if tm.isPieceLocationValid(updatedPiece) {
		tm.currentPiece = updatedPiece

		tm.invalidateDrawableCellsCache()
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
	if currentPiece == nil {
		return
	}

	for _, coordinate := range currentPiece.Coordinates() {
		stackCell := &tm.stackCells[coordinate.Row()][coordinate.Column()]
		stackCell.occupied = true
		stackCell.color = currentPiece.Color()
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

func (tm *tetrisModel) UpdateDuration() time.Duration {
	switch {
	case tm.lines < 10 || tm.gameOver:
		return 500 * time.Millisecond
	case tm.lines < 20:
		return 450 * time.Millisecond
	case tm.lines < 30:
		return 400 * time.Millisecond
	case tm.lines < 40:
		return 350 * time.Millisecond
	case tm.lines < 50:
		return 300 * time.Millisecond
	case tm.lines < 60:
		return 250 * time.Millisecond
	case tm.lines < 70:
		return 200 * time.Millisecond
	default:
		return 150 * time.Millisecond
	}
}

type drawableCellsCache struct {
	drawableCellsMap DrawableCellsMap
	valid            bool
}

func (cache *drawableCellsCache) invalidate() {
	cache.valid = false
}

func (cache *drawableCellsCache) drawableCells(tm *tetrisModel) DrawableCellsMap {
	if cache.valid {
		return cache.drawableCellsMap
	}

	if cache.drawableCellsMap == nil {
		cache.drawableCellsMap = make(DrawableCellsMap, coordinate.BoardModelNumCells)
	}

	clear(cache.drawableCellsMap)

	for row := range tm.stackCells {
		for column := range tm.stackCells[row] {
			stackCell := &tm.stackCells[row][column]
			if stackCell.occupied {
				coord := coordinate.NewTetrisModelCoordinate(row, column)
				cache.drawableCellsMap[coord] = stackCell.color
			}
		}
	}

	if tm.currentPiece != nil {
		for _, coord := range tm.currentPiece.Coordinates() {
			cache.drawableCellsMap[coord] = tm.currentPiece.Color()
		}
	}

	cache.valid = true

	return cache.drawableCellsMap
}
