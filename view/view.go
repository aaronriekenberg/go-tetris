package view

import (
	"fmt"
	"os"

	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/aaronriekenberg/go-tetris/model"
	"github.com/gdamore/tcell/v2"
)

type View struct {
	screen             tcell.Screen
	drawableCellsModel model.DrawableCellsModel
}

func NewView(
	drawableCellsModel model.DrawableCellsModel,
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
		screen:             screen,
		drawableCellsModel: drawableCellsModel,
	}
}

var bgStyle = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

func (view *View) Draw() {

	drawableCells := view.drawableCellsModel.DrawableCells()

	w, h := view.screen.Size()

	// s.Clear()

	if w < (common.BoardColumns*2) || h < common.BoardRows {
		view.screen.Show()
		return
	}

	boardLeftX := (w - (common.BoardColumns * 2)) / 2
	boardTopY := (h - common.BoardRows) / 2

	for viewColumn := 0; viewColumn < (common.BoardColumns * 2); viewColumn += 2 {
		for viewRow := 0; viewRow < (common.BoardRows); viewRow += 1 {
			var comb []rune
			modelRow := viewRow
			modelColumn := (viewColumn / 2)

			modelCell := drawableCells[modelRow][modelColumn]
			if modelCell.Occupied() {
				fgStyle := tcell.StyleDefault.
					Foreground(modelCell.Color()).
					Background(modelCell.Color())

				view.screen.SetContent(boardLeftX+viewColumn, boardTopY+viewRow, ' ', comb, fgStyle)
				view.screen.SetContent(boardLeftX+viewColumn+1, boardTopY+viewRow, ' ', comb, fgStyle)
			} else {
				view.screen.SetContent(boardLeftX+viewColumn, boardTopY+viewRow, ' ', comb, bgStyle)
				view.screen.SetContent(boardLeftX+viewColumn+1, boardTopY+viewRow, ' ', comb, bgStyle)
			}
		}
	}

	view.screen.Show()
}

func (view *View) HandleResizeEvent() {
	view.screen.Clear()
	view.screen.Sync()
	view.Draw()
}

func (view *View) Finalize() {
	view.screen.Fini()
}

func (view *View) Screen() tcell.Screen {
	return view.screen
}

// func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
// 	for _, c := range str {
// 		var comb []rune
// 		w := runewidth.RuneWidth(c)
// 		if w == 0 {
// 			comb = []rune{c}
// 			c = ' '
// 			w = 1
// 		}
// 		s.SetContent(x, y, c, comb, style)
// 		x += w
// 	}
// }
