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

func (view *view) viewBoardCoordinates() (result viewBoardCoordinates, ok bool) {
	w, h := view.tcellScreen.Size()

	if w < boardWidthViewCells || h < boardHeightViewCells {
		ok = false
		return
	}

	ok = true

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

	unoccupiedCellStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

	for viewColumn := 0; viewColumn < boardWidthViewCells; viewColumn += 2 {
		for viewRow := 0; viewRow < boardHeightViewCells; viewRow += 1 {
			var comb []rune
			modelCoordinate := coordinate.NewTetrisModelCoordinate(
				viewRow, (viewColumn / 2),
			)

			cellColor := drawableCells[modelCoordinate]

			style := unoccupiedCellStyle
			if cellColor.Valid() {
				style = tcell.StyleDefault.Foreground(cellColor).Background(cellColor)
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

	viewBoardCoordinates, ok := view.viewBoardCoordinates()
	if !ok {
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
	viewBoardCoordinates, ok := view.viewBoardCoordinates()
	if !ok {
		return
	}

	switch {
	case y < viewBoardCoordinates.boardTopY:
		view.model.RotateCurrentPiece()

	case y > viewBoardCoordinates.boardBottomY:
		if eventTime.Sub(view.lastMoveDownButtonEventTime) <= 200*time.Millisecond {
			// double click
			view.model.DropCurrentPiece()
		} else {
			view.model.MoveCurrentPieceDown()
		}
		view.lastMoveDownButtonEventTime = eventTime

	case utils.IntegerAbs(x-viewBoardCoordinates.boardLeftX) < utils.IntegerAbs(x-viewBoardCoordinates.boardRightX):
		view.model.MoveCurrentPieceLeft()

	default:
		view.model.MoveCurrentPieceRight()
	}

	view.Draw()
}
