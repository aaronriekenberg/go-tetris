package main

import (
	"fmt"
	"os"
	"slices"
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
	color    tcell.Color
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
	stackCells := make([][]tetrisModelCell, common.BoardRows)

	for row := 0; row < common.BoardRows; row += 1 {
		stackCells[row] = make([]tetrisModelCell, common.BoardColumns)
	}

	tetrisModel.stackCells = stackCells
}

func (tetrisModel *tetrisModel) recomputeDrawableCells() {
	drawableCells := make([][]tetrisModelCell, common.BoardRows)

	for row := 0; row < common.BoardRows; row += 1 {
		drawableCells[row] = make([]tetrisModelCell, common.BoardColumns)
		for column := 0; column < common.BoardColumns; column += 1 {
			drawableCells[row][column] = tetrisModel.stackCells[row][column]
		}
	}

	if tetrisModel.currentPiece != nil {
		for _, coordinates := range tetrisModel.currentPiece.Coordinates() {
			drawableCells[coordinates.Row()][coordinates.Column()].occupied = true
			drawableCells[coordinates.Row()][coordinates.Column()].color = tetrisModel.currentPiece.Color()
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
		} else if tetrisModel.stackCells[coordinate.Row()][coordinate.Column()].occupied {
			return false
		}
	}
	return true
}

func (tetrisModel *tetrisModel) addNewPiece() {
	centerCoordinate := common.NewTetrisModelCoordinate(
		0,
		(common.BoardColumns/2)-1,
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

		newCenterCoordinate := centerCoordinate.AddRows(1)

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

		newCenterCoordinate := centerCoordinate.AddColumns(-1)

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

		newCenterCoordinate := centerCoordinate.AddColumns(1)

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
			tetrisModel.stackCells[coordinate.Row()][coordinate.Column()].occupied = true
			tetrisModel.stackCells[coordinate.Row()][coordinate.Column()].color = currentPiece.Color()
		}
	}
	tetrisModel.currentPiece = nil

	tetrisModel.handleFilledStackRows()
}

func (tetrisModel *tetrisModel) handleFilledStackRows() {
	row := common.BoardRows - 1

	for row >= 0 {
		rowIsFull := true
		for _, cell := range tetrisModel.stackCells[row] {
			if !cell.occupied {
				rowIsFull = false
				break
			}
		}
		if rowIsFull {
			tetrisModel.stackCells = slices.Delete(tetrisModel.stackCells, row, row+1)
			tetrisModel.stackCells = slices.Insert(tetrisModel.stackCells, 0, make([]tetrisModelCell, common.BoardColumns))
			// numLines += 1
		} else {
			row -= 1
		}
	}
}

func (tetrisModel *tetrisModel) periodicUpdate() {
	if tetrisModel.currentPiece == nil {
		tetrisModel.addNewPiece()
	} else {
		tetrisModel.moveCurrentPieceDown()
	}
}

var bgStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorWhite)

func drawBoard(
	tetrisModel *tetrisModel,
	s tcell.Screen,
) {
	// startTime := time.Now()

	tetrisModel.recomputeDrawableCells()

	w, h := s.Size()

	// s.Clear()

	if w < (common.BoardColumns*2) || h < common.BoardRows {
		s.Show()
		return
	}

	boardLeftX := (w - (common.BoardColumns * 2)) / 2
	boardTopY := (h - common.BoardRows) / 2

	for column := 0; column < (common.BoardColumns * 2); column += 2 {
		for row := 0; row < (common.BoardRows); row += 1 {
			var comb []rune
			modelRow := row
			modelColumn := (column / 2)
			if tetrisModel.drawableCells[modelRow][modelColumn].occupied {
				fgStyle := tcell.StyleDefault.
					Foreground(tetrisModel.drawableCells[modelRow][modelColumn].color).
					Background(tetrisModel.drawableCells[modelRow][modelColumn].color)

				s.SetContent(boardLeftX+column, boardTopY+row, ' ', comb, fgStyle)
				s.SetContent(boardLeftX+column+1, boardTopY+row, ' ', comb, fgStyle)
			} else {
				s.SetContent(boardLeftX+column, boardTopY+row, ' ', comb, bgStyle)
				s.SetContent(boardLeftX+column+1, boardTopY+row, ' ', comb, bgStyle)
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
