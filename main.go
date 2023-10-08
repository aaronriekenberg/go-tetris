package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"

	"github.com/mattn/go-runewidth"
)

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
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

// func displayHelloWorld(s tcell.Screen) {
// 	w, h := s.Size()
// 	s.Clear()
// 	style := tcell.StyleDefault.Foreground(tcell.ColorBlue.TrueColor()).Background(tcell.ColorWhite)
// 	emitStr(s, w/2-12, h/2, style, fmt.Sprintf("width = %03v height = %03v", w, h))
// 	emitStr(s, w/2-12, h/2+1, style, "                            ")
// 	s.Show()
// }

const (
	boardWidth  = 12
	boardHeight = 16
)

// (0, 0) is topmost and leftmost cell
type tetrisModelCoordinate struct {
	x int
	y int
}

func (tmc tetrisModelCoordinate) valid() bool {
	return (tmc.x >= 0) && (tmc.x < boardWidth) && (tmc.y >= 0) && (tmc.y < boardHeight)
}

func (tmc tetrisModelCoordinate) addY(delta int) tetrisModelCoordinate {
	tmc.y += delta
	return tmc
}

func (tmc tetrisModelCoordinate) addX(delta int) tetrisModelCoordinate {
	tmc.x += delta
	return tmc
}

type tetrisPieceType int

const (
	squarePieceType tetrisPieceType = iota
)

type tetrisPiece struct {
	pieceType        tetrisPieceType
	centerCoordinate tetrisModelCoordinate
}

func (tetrisPiece tetrisPiece) cloneWithNewCenterCoordinate(
	newCenterCoordinate tetrisModelCoordinate,
) tetrisPiece {
	tetrisPiece.centerCoordinate = newCenterCoordinate
	return tetrisPiece
}

func (tetrisPiece tetrisPiece) coordinates() []tetrisModelCoordinate {
	switch tetrisPiece.pieceType {
	case squarePieceType:
		return []tetrisModelCoordinate{
			{x: tetrisPiece.centerCoordinate.x, y: tetrisPiece.centerCoordinate.y},
			{x: tetrisPiece.centerCoordinate.x + 1, y: tetrisPiece.centerCoordinate.y},
			{x: tetrisPiece.centerCoordinate.x, y: tetrisPiece.centerCoordinate.y + 1},
			{x: tetrisPiece.centerCoordinate.x + 1, y: tetrisPiece.centerCoordinate.y + 1},
		}
	}

	return nil
}

type tetrisModelCell struct {
	occupied bool
}

type tetrisModel struct {
	drawableCells [][]tetrisModelCell
	currentPiece  *tetrisPiece
	stackCells    [][]tetrisModelCell
}

func newTetrisModel() *tetrisModel {
	tetrisModel := &tetrisModel{}

	tetrisModel.initializeStackCells()

	tetrisModel.recomputeDrawableCells()

	return tetrisModel
}

func (tetrisModel *tetrisModel) initializeStackCells() {
	stackCells := make([][]tetrisModelCell, boardWidth)

	for x := 0; x < boardWidth; x += 1 {
		stackCells[x] = make([]tetrisModelCell, boardHeight)
	}

	tetrisModel.stackCells = stackCells
}

func (tetrisModel *tetrisModel) recomputeDrawableCells() {
	drawableCells := make([][]tetrisModelCell, boardWidth)

	for x := 0; x < boardWidth; x += 1 {
		drawableCells[x] = make([]tetrisModelCell, boardHeight)
		for y := 0; y < boardHeight; y += 1 {
			drawableCells[x][y] = tetrisModel.stackCells[x][y]
		}
	}

	if tetrisModel.currentPiece != nil {
		for _, coordinates := range tetrisModel.currentPiece.coordinates() {
			drawableCells[coordinates.x][coordinates.y].occupied = true
		}
	}

	tetrisModel.drawableCells = drawableCells
}

func (tetrisModel *tetrisModel) isPieceLocationValid(
	tetrisPiece tetrisPiece,
) bool {
	for _, coordinate := range tetrisPiece.coordinates() {
		if !coordinate.valid() {
			return false
		} else if tetrisModel.stackCells[coordinate.x][coordinate.y].occupied {
			return false
		}
	}
	return true
}

func (tetrisModel *tetrisModel) addNewPiece() {
	centerCoordinate := tetrisModelCoordinate{
		x: (boardWidth / 2) - 1,
		y: 0,
	}

	newPiece := tetrisPiece{
		pieceType:        squarePieceType,
		centerCoordinate: centerCoordinate,
	}

	if !tetrisModel.isPieceLocationValid(newPiece) {
		fmt.Printf("unable to add newPiece")
		return
	}

	tetrisModel.currentPiece = &newPiece
}

func (tetrisModel *tetrisModel) moveCurrentPieceDown() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.centerCoordinate

		newCenterCoordinate := centerCoordinate.addY(1)

		updatedPiece := currentPiece.cloneWithNewCenterCoordinate(newCenterCoordinate)

		if !tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.addCurrentPieceToStack()
		} else {
			tetrisModel.currentPiece = &updatedPiece
		}
	}
}

func (tetrisModel *tetrisModel) moveCurrentPieceLeft() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.centerCoordinate

		newCenterCoordinate := centerCoordinate.addX(-1)

		updatedPiece := currentPiece.cloneWithNewCenterCoordinate(newCenterCoordinate)

		if tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.currentPiece = &updatedPiece
		}
	}
}

func (tetrisModel *tetrisModel) moveCurrentPieceRight() {
	currentPiece := tetrisModel.currentPiece
	if currentPiece != nil {
		centerCoordinate := currentPiece.centerCoordinate

		newCenterCoordinate := centerCoordinate.addX(1)

		updatedPiece := currentPiece.cloneWithNewCenterCoordinate(newCenterCoordinate)

		if tetrisModel.isPieceLocationValid(updatedPiece) {
			tetrisModel.currentPiece = &updatedPiece
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
		for _, coordinate := range currentPiece.coordinates() {
			tetrisModel.stackCells[coordinate.x][coordinate.y].occupied = true
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

	tetrisModel.recomputeDrawableCells()
}

var fgStyle = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
var bgStyle = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)

func drawBoard(
	tetrisModel *tetrisModel,
	s tcell.Screen,
) {
	// startTime := time.Now()

	w, h := s.Size()

	// s.Clear()

	if w < boardWidth || h < boardHeight {
		s.Show()
		return
	}

	boardLeftX := (w - (boardWidth * 2)) / 2
	boardTopY := (h - boardHeight) / 2

	for x := 0; x < (boardWidth * 2); x += 2 {
		for y := 0; y < (boardHeight); y += 1 {
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

	// s.Sync()

	// delta := time.Since(startTime)

	// fmt.Printf("delta = %v", delta)

}

// func makebox(s tcell.Screen) {
// 	w, h := s.Size()

// 	s.Clear()

// 	// if w < 300 || h < 550 {
// 	// 	s.Show()
// 	// 	return
// 	// }
// 	st := tcell.StyleDefault
// 	gl := 'A'

// 	s.SetCell((w-width)/2, (h-height)/2, st, gl)

// 	// glyphs := []rune{'@', '#', '&', '*', '=', '%', 'Z', 'A'}

// 	// lx := rand.Int() % w
// 	// ly := rand.Int() % h
// 	// lw := rand.Int() % (w - lx)
// 	// lh := rand.Int() % (h - ly)
// 	// st := tcell.StyleDefault
// 	// gl := ' '
// 	// if s.Colors() > 256 {
// 	// 	rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
// 	// 	st = st.Background(rgb)
// 	// } else if s.Colors() > 1 {
// 	// 	st = st.Background(tcell.Color(rand.Int()%s.Colors()) | tcell.ColorValid)
// 	// } else {
// 	// 	st = st.Reverse(rand.Int()%2 == 0)
// 	// 	gl = glyphs[rand.Int()%len(glyphs)]
// 	// }
// 	rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
// 	st = st.Background(rgb)

// 	lx := rand.Int() % w
// 	ly := rand.Int() % h
// 	lw := rand.Int() % (w - lx)
// 	lh := rand.Int() % (h - ly)
// 	for row := 0; row < lh; row++ {
// 		for col := 0; col < lw; col++ {
// 			s.SetCell(lx+col, ly+row, st, gl)
// 		}
// 	}
// 	s.Show()
// }

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

	// eventChannel := make(chan tcell.Event, 1000)
	// quitChannel := make(chan struct{}, 10)
	// go s.ChannelEvents(eventChannel, quitChannel)

	periodicUpdateTicker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			<-periodicUpdateTicker.C

			// fmt.Printf("calling postEvent")
			s.PostEvent(tcell.NewEventInterrupt(periodicUpdateInterruptCustomEvent{}))
		}
	}()

	// redrawTicker := time.NewTicker(50 * time.Millisecond)
	// go func() {
	// 	for {
	// 		<-redrawTicker.C

	// 		// fmt.Printf("calling postEvent")
	// 		s.PostEvent(tcell.NewEventInterrupt(redrawInterruptCustomEvent{}))
	// 	}
	// }()

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
		// select {
		// case ev := <-eventChannel:
		ev := s.PollEvent()
		// fmt.Printf("got event")
		switch ev := ev.(type) {
		case *tcell.EventInterrupt:
			switch ev.Data().(type) {
			case periodicUpdateInterruptCustomEvent:
				tetrisModel.periodicUpdate()
				drawBoard(tetrisModel, s)

				// case redrawInterruptCustomEvent:
				// 	drawBoard(tetrisModel, s)
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
			fmt.Printf("EvetnResize")
			s.Clear()
			s.Sync()
			drawBoard(tetrisModel, s)
		}
	}

	fmt.Printf("end main")

}

type periodicUpdateInterruptCustomEvent struct{}

// type redrawInterruptCustomEvent struct{}
