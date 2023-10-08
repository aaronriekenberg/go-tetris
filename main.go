package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aaronriekenberg/go-tetris/common"
	"github.com/aaronriekenberg/go-tetris/pieces"

	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"
)

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

type tetrisModelCell struct {
	occupied bool
}

type tetrisModel struct {
	drawableCells [][]tetrisModelCell
	currentPiece  pieces.TetrisPiece
	stackCells    [][]tetrisModelCell
}

func newTetrisModel() *tetrisModel {
	tetrisModel := &tetrisModel{}

	tetrisModel.initializeStackCells()

	tetrisModel.recomputeDrawableCells()

	return tetrisModel
}

func (tetrisModel *tetrisModel) initializeStackCells() {
	stackCells := make([][]tetrisModelCell, common.BoardWidth)

	for x := 0; x < common.BoardWidth; x += 1 {
		stackCells[x] = make([]tetrisModelCell, common.BoardHeight)
	}

	tetrisModel.stackCells = stackCells
}

func (tetrisModel *tetrisModel) recomputeDrawableCells() {
	drawableCells := make([][]tetrisModelCell, common.BoardWidth)

	for x := 0; x < common.BoardWidth; x += 1 {
		drawableCells[x] = make([]tetrisModelCell, common.BoardHeight)
		for y := 0; y < common.BoardHeight; y += 1 {
			drawableCells[x][y] = tetrisModel.stackCells[x][y]
		}
	}

	if tetrisModel.currentPiece != nil {
		for _, coordinates := range tetrisModel.currentPiece.Coordinates() {
			drawableCells[coordinates.X()][coordinates.Y()].occupied = true
		}
	}

	tetrisModel.drawableCells = drawableCells
}

func (tetrisModel *tetrisModel) isPieceLocationValid(
	tetrisPiece pieces.TetrisPiece,
) bool {
	for _, coordinate := range tetrisPiece.Coordinates() {
		if !coordinate.Valid() {
			return false
		} else if tetrisModel.stackCells[coordinate.X()][coordinate.Y()].occupied {
			return false
		}
	}
	return true
}

func (tetrisModel *tetrisModel) addNewPiece() {
	centerCoordinate := common.NewTetrisModelCoordinate(
		(common.BoardWidth/2)-1,
		0,
	)

	newPiece := pieces.CreateRandomPiece(centerCoordinate)

	if !tetrisModel.isPieceLocationValid(newPiece) {
		fmt.Printf("unable to add newPiece")
		return
	}

	tetrisModel.currentPiece = newPiece
}

func (tetrisModel *tetrisModel) moveCurrentPieceDown() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddY(1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if !tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.addCurrentPieceToStack()
		} else {
			tetrisModel.currentPiece = updatedPiece
		}
	}
}

func (tetrisModel *tetrisModel) moveCurrentPieceLeft() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddX(-1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.currentPiece = updatedPiece
		}
	}
}

func (tetrisModel *tetrisModel) moveCurrentPieceRight() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.CenterCoordinate()

		newCenterCoordinate := centerCoordinate.AddX(1)

		updatedPiece := currentPiece.CloneWithNewCenterCoordinate(newCenterCoordinate)

		if tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.currentPiece = updatedPiece
		}
	}
}

func (tetrisModel *tetrisModel) dropCurrentPiece() {
	for tetrisModel.currentPiece != nil {
		tetrisModel.moveCurrentPieceDown()
	}
}

func (tetrisModel *tetrisModel) addCurrentPieceToStack() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		for _, coordinate := range currentPiece.Coordinates() {
			tetrisModel.stackCells[coordinate.X()][coordinate.Y()].occupied = true
		}
	}
	tetrisModel.currentPiece = nil
}

func (tetrisModel *tetrisModel) periodicUpdate() {
	if tetrisModel.currentPiece == nil {
		tetrisModel.addNewPiece()
	} else {
		tetrisModel.moveCurrentPieceDown()
	}
}

var fgStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
var bgStyle = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)

func drawBoard(
	tetrisModel *tetrisModel,
	s tcell.Screen,
) {
	// startTime := time.Now()

	tetrisModel.recomputeDrawableCells()

	w, h := s.Size()

	// s.Clear()

	if w < common.BoardWidth || h < common.BoardHeight {
		s.Show()
		return
	}

	boardLeftX := (w - (common.BoardWidth * 2)) / 2
	boardTopY := (h - common.BoardHeight) / 2

	for x := 0; x < (common.BoardWidth * 2); x += 2 {
		for y := 0; y < (common.BoardHeight); y += 1 {
			var comb []rune
			if tetrisModel.drawableCells[x/2][y].occupied {
				s.SetContent(boardLeftX+x, boardTopY+y, ' ', comb, fgStyle)
				s.SetContent(boardLeftX+x+1, boardTopY+y, ' ', comb, fgStyle)
			} else {
				s.SetContent(boardLeftX+x, boardTopY+y, ' ', comb, bgStyle)
				s.SetContent(boardLeftX+x+1, boardTopY+y, ' ', comb, bgStyle)
			}
		}
	}

	s.Show()

}

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.Clear()

	tetrisModel := newTetrisModel()

	periodicUpdateTicker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			<-periodicUpdateTicker.C

			s.PostEvent(tcell.NewEventInterrupt(periodicUpdateInterruptCustomEvent{}))
		}
	}()

	done := false

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.

		done = true

		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	for !done {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventInterrupt:
			switch ev.Data().(type) {
			case periodicUpdateInterruptCustomEvent:
				tetrisModel.periodicUpdate()
				drawBoard(tetrisModel, s)
			}
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				quit()
			case tcell.KeyLeft:
				tetrisModel.moveCurrentPieceLeft()
				drawBoard(tetrisModel, s)
			case tcell.KeyRight:
				tetrisModel.moveCurrentPieceRight()
				drawBoard(tetrisModel, s)
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'q':
					quit()
				case ' ':
					tetrisModel.dropCurrentPiece()
					drawBoard(tetrisModel, s)
				}
			}
		case *tcell.EventResize:
			s.Clear()
			s.Sync()
			drawBoard(tetrisModel, s)
		}
	}

}

type periodicUpdateInterruptCustomEvent struct{}
