package view

import (
	"fmt"
	"os"

	"github.com/aaronriekenberg/go-tetris/coord"
	"github.com/aaronriekenberg/go-tetris/model"
	"github.com/aaronriekenberg/go-tetris/version"

	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"

	"github.com/mattn/go-runewidth"
)

type View struct {
	screen            tcell.Screen
	drawableInfoModel model.DrawableInfoModel
	showVersion       bool
}

func NewView(
	drawableInfoModel model.DrawableInfoModel,
) *View {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	screen.Clear()

	return &View{
		screen:            screen,
		drawableInfoModel: drawableInfoModel,
	}
}

func (view *View) Clear() {
	view.screen.Clear()
}

const boardWidthCells = coord.BoardColumns * 2
const boardHeightCells = coord.BoardRows

func (view *View) drawBoard(
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

			view.screen.SetContent(x, y, ' ', comb, style)
			view.screen.SetContent(x+1, y, ' ', comb, style)
		}
	}
}

func (view *View) drawTextFields(
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

func (view *View) Draw() {

	w, h := view.screen.Size()

	if w < boardWidthCells || h < boardHeightCells {
		view.screen.Clear()
		view.screen.Show()
		return
	}

	boardLeftX := (w - boardWidthCells) / 2

	boardTopY := (h - boardHeightCells) / 2

	view.drawBoard(boardLeftX, boardTopY)

	view.drawTextFields(boardLeftX, boardTopY)

	view.screen.Show()
}

func (view *View) emitStr(x, y int, style tcell.Style, str string) {
	s := view.screen

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

func (view *View) HandleResizeEvent() {
	view.screen.Clear()
	view.screen.Sync()
	view.Draw()
}

func (view *View) PollEvent() tcell.Event {
	return view.screen.PollEvent()
}

func (view *View) PostEvent(ev tcell.Event) error {
	return view.screen.PostEvent(ev)
}

func (view *View) Finalize() {
	view.screen.Fini()
}

func (view *View) ToggleShowVersion() {
	view.showVersion = !view.showVersion
}
