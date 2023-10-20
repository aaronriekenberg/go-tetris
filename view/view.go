package view

import (
	"fmt"

	"github.com/aaronriekenberg/go-tetris/coordinate"
	"github.com/aaronriekenberg/go-tetris/model"
	"github.com/aaronriekenberg/go-tetris/version"

	"github.com/gdamore/tcell/v2"

	"github.com/mattn/go-runewidth"
)

type View interface {
	Clear()
	Draw()
	HandleResizeEvent()
	ToggleShowVersion()
}

type view struct {
	tcellScreen       tcell.Screen
	drawableInfoModel model.DrawableInfoModel
	showVersion       bool
}

func NewView(
	screen tcell.Screen,
	drawableInfoModel model.DrawableInfoModel,
) View {

	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	screen.Clear()

	return &view{
		tcellScreen:       screen,
		drawableInfoModel: drawableInfoModel,
	}
}

func (view *view) Clear() {
	view.tcellScreen.Clear()
}

const boardWidthCells = coordinate.BoardColumns * 2
const boardHeightCells = coordinate.BoardRows

func (view *view) drawBoard(
	boardLeftX, boardTopY int,
) {
	drawableCells := view.drawableInfoModel.DrawableCells()

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
		fmt.Sprintf("Lines: % 3v", view.drawableInfoModel.Lines()),
	)

	if view.drawableInfoModel.GameOver() {
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

	w, h := view.tcellScreen.Size()

	if w < boardWidthCells || h < boardHeightCells {
		view.tcellScreen.Clear()
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
