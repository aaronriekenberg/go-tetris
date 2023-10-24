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
	boardWidthCells  = coordinate.BoardColumns * 2
	boardHeightCells = coordinate.BoardRows
)

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

func (view *view) drawBoard(
	boardLeftX, boardTopY int,
) {
	drawableCells := view.model.DrawableCells()

	bgStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

	for viewColumn := 0; viewColumn < boardWidthCells; viewColumn += 2 {
		for viewRow := 0; viewRow < boardHeightCells; viewRow += 1 {
			var comb []rune
			modelRow := viewRow
			modelColumn := (viewColumn / 2)

			modelCell := drawableCells[modelRow][modelColumn]

			style := bgStyle
			if modelCell.Occupied() {
				style = tcell.StyleDefault.Foreground(modelCell.Color()).Background(modelCell.Color())
			}

			x, y := boardLeftX+viewColumn, boardTopY+viewRow

			view.tcellScreen.SetContent(x, y, ' ', comb, style)
			view.tcellScreen.SetContent(x+1, y, ' ', comb, style)
		}
	}
}

func (view *view) drawTextFields(
	boardLeftX, boardTopY int,
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
		boardLeftX+(boardWidthCells/2)-5,
		boardTopY+boardHeightCells+1,
		textStyle,
		fmt.Sprintf("Lines: % 3v", view.model.Lines()),
	)

	if view.model.GameOver() {
		textStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

		view.emitStr(
			boardLeftX+(boardWidthCells/2)-4,
			boardTopY+boardHeightCells+3,
			textStyle,
			"GAME OVER",
		)
	}

}

func (view *view) Draw() {

	view.tcellScreen.Clear()

	w, h := view.tcellScreen.Size()

	if w < boardWidthCells || h < boardHeightCells {
		view.tcellScreen.Show()
		return
	}

	boardLeftX := (w - boardWidthCells) / 2

	boardTopY := (h - boardHeightCells) / 2

	view.drawBoard(boardLeftX, boardTopY)

	view.drawTextFields(boardLeftX, boardTopY)

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
}

func (view *view) HandleButton1PressEvent(
	x, y int,
	eventTime time.Time,
) {
	w, h := view.tcellScreen.Size()

	if w < boardWidthCells || h < boardHeightCells {
		return
	}

	boardLeftX := (w - boardWidthCells) / 2

	boardRightX := boardLeftX + boardWidthCells

	boardTopY := (h - boardHeightCells) / 2

	boardBottomY := boardTopY + boardHeightCells

	if y < boardTopY {
		view.model.RotateCurrentPiece()
		view.Draw()
		return
	}

	if y > boardBottomY {
		if time.Since(view.lastMoveDownButtonEventTime) <= 200*time.Millisecond {
			// double click
			view.model.DropCurrentPiece()
		} else {
			view.model.MoveCurrentPieceDown()
		}
		view.Draw()
		view.lastMoveDownButtonEventTime = eventTime
		return
	}

	if utils.IntegerAbs(x-boardLeftX) < utils.IntegerAbs(x-boardRightX) {
		view.model.MoveCurrentPieceLeft()
		view.Draw()
		return
	} else {
		view.model.MoveCurrentPieceRight()
		view.Draw()
		return
	}

}
