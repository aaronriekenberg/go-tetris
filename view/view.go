package view

import (
	"fmt"
	"os"

	"github.com/aaronriekenberg/go-tetris/coord"
	"github.com/aaronriekenberg/go-tetris/model"

	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"

	"github.com/mattn/go-runewidth"
)

type View struct {
	screen            tcell.Screen
	drawableInfoModel model.DrawableInfoModel
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

var bgStyle = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

func (view *View) Draw() {

	drawableCells := view.drawableInfoModel.DrawableCells()

	w, h := view.screen.Size()

	if w < (coord.BoardColumns*2) || h < coord.BoardRows {
		view.screen.Clear()
		view.screen.Show()
		return
	}

	const boardWidthCells = coord.BoardColumns * 2
	boardLeftX := (w - boardWidthCells) / 2

	const boardHeightCells = coord.BoardRows
	boardTopY := (h - boardHeightCells) / 2

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

			view.screen.SetContent(boardLeftX+viewColumn, boardTopY+viewRow, ' ', comb, style)
			view.screen.SetContent(boardLeftX+viewColumn+1, boardTopY+viewRow, ' ', comb, style)
		}
	}

	textStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)

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

	view.screen.Show()
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
