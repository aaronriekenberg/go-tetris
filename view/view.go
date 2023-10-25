package view

import (
	"fmt"
	"time"

	"github.com/aaronriekenberg/go-tetris/coordinate"
	"github.com/aaronriekenberg/go-tetris/model"
	"github.com/aaronriekenberg/go-tetris/utils"
	"github.com/aaronriekenberg/go-tetris/version"

	"github.com/gdamore/tcell/v2"

	"github.com/mattn/go-runewidth"
)

const (
	boardWidthViewCells  = coordinate.BoardModelColumns * 2
	boardHeightViewCells = coordinate.BoardModelRows
)

type viewBoardCoordinates struct {
	valid        bool
	boardLeftX   int
	boardRightX  int
	boardTopY    int
	boardBottomY int
}

type View interface {
	Draw()
	HandleResizeEvent()
	ToggleShowVersion()
	HandleButton1PressEvent(
		x, y int,
		eventTime time.Time,
	)
}

type view struct {
	tcellScreen                 tcell.Screen
	model                       model.TetrisModel
	showVersion                 bool
	lastMoveDownButtonEventTime time.Time
}

func NewView(
	screen tcell.Screen,
	model model.TetrisModel,
) View {

	if utils.RunningInWASM() {
		screen.SetSize(40, 34)
	}

	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	screen.Clear()

	return &view{
		tcellScreen: screen,
		model:       model,
	}
}

func (view *view) viewBoardCoordinates() (result viewBoardCoordinates) {
	w, h := view.tcellScreen.Size()

	if w < boardWidthViewCells || h < boardHeightViewCells {
		result.valid = false
		return
	}

	result.valid = true

	result.boardLeftX = (w - boardWidthViewCells) / 2

	result.boardRightX = result.boardLeftX + boardWidthViewCells - 2

	result.boardTopY = (h - boardHeightViewCells) / 2

	result.boardBottomY = result.boardTopY + boardHeightViewCells - 1

	return
}

func (view *view) drawBoard(
	viewBoardCoordinates viewBoardCoordinates,
) {
	drawableCells := view.model.DrawableCells()

	bgStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

	for viewColumn := 0; viewColumn < boardWidthViewCells; viewColumn += 2 {
		for viewRow := 0; viewRow < boardHeightViewCells; viewRow += 1 {
			var comb []rune
			modelRow := viewRow
			modelColumn := (viewColumn / 2)

			modelCell := drawableCells[modelRow][modelColumn]

			style := bgStyle
			if modelCell.Occupied() {
				style = tcell.StyleDefault.Foreground(modelCell.Color()).Background(modelCell.Color())
			}

			x, y := viewBoardCoordinates.boardLeftX+viewColumn, viewBoardCoordinates.boardTopY+viewRow

			view.tcellScreen.SetContent(x, y, ' ', comb, style)
			view.tcellScreen.SetContent(x+1, y, ' ', comb, style)
		}
	}
}

func (view *view) drawTextFields(
	viewBoardCoordinates viewBoardCoordinates,
) {
	textStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)

	if view.showVersion {
		view.emitStr(
			0, 0,
			textStyle,
			version.VersionInfoString(),
		)
	}

	view.emitStr(
		viewBoardCoordinates.boardLeftX+(boardWidthViewCells/2)-5,
		viewBoardCoordinates.boardTopY+boardHeightViewCells+1,
		textStyle,
		fmt.Sprintf("Lines: % 3v", view.model.Lines()),
	)

	if view.model.GameOver() {
		textStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

		view.emitStr(
			viewBoardCoordinates.boardLeftX+(boardWidthViewCells/2)-4,
			viewBoardCoordinates.boardTopY+boardHeightViewCells+3,
			textStyle,
			"GAME OVER",
		)
	}

}

func (view *view) Draw() {

	view.tcellScreen.Clear()

	viewBoardCoordinates := view.viewBoardCoordinates()
	if !viewBoardCoordinates.valid {
		return
	}

	view.drawBoard(viewBoardCoordinates)

	view.drawTextFields(viewBoardCoordinates)

	view.tcellScreen.Show()
}

func (view *view) emitStr(x, y int, style tcell.Style, str string) {
	s := view.tcellScreen

	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func (view *view) HandleResizeEvent() {
	view.tcellScreen.Clear()
	view.tcellScreen.Sync()
	view.Draw()
}

func (view *view) ToggleShowVersion() {
	view.showVersion = !view.showVersion

	view.Draw()
}

func (view *view) HandleButton1PressEvent(
	x, y int,
	eventTime time.Time,
) {
	viewBoardCoordinates := view.viewBoardCoordinates()
	if !viewBoardCoordinates.valid {
		return
	}

	switch {
	case y < viewBoardCoordinates.boardTopY:
		view.model.RotateCurrentPiece()

	case y > viewBoardCoordinates.boardBottomY:
		if time.Since(view.lastMoveDownButtonEventTime) <= 200*time.Millisecond {
			// double click
			view.model.DropCurrentPiece()
		} else {
			view.model.MoveCurrentPieceDown()
		}
		view.lastMoveDownButtonEventTime = eventTime

	case utils.Abs(x-viewBoardCoordinates.boardLeftX) < utils.Abs(x-viewBoardCoordinates.boardRightX):
		view.model.MoveCurrentPieceLeft()

	default:
		view.model.MoveCurrentPieceRight()
	}

	view.Draw()

}
