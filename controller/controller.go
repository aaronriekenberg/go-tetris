package controller

import (
	"runtime"
	"time"

	"github.com/aaronriekenberg/go-tetris/model"
	"github.com/aaronriekenberg/go-tetris/view"

	"github.com/gdamore/tcell/v2"
)

func Run() {
	tetrisModel := model.NewTetrisModel()

	screen := view.NewScreen()

	view := view.NewView(
		screen,
		tetrisModel,
	)

	runEventLoop(
		screen,
		view,
		tetrisModel,
	)
}

type mouseEventInfo struct {
	x int
	y int
}

func runEventLoop(
	eventSource view.ScreenEventSource,
	view view.View,
	tetrisModel model.TetrisModel,
) {
	runningInWASM := runtime.GOARCH == "wasm"

	periodicUpdateTicker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			<-periodicUpdateTicker.C

			eventSource.PostEvent(tcell.NewEventInterrupt(periodicUpdateInterruptCustomEvent{}))
		}
	}()

	var repeatingMouseEvent *mouseEventInfo

	done := false

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.

		done = true

		maybePanic := recover()
		eventSource.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	lastModelUpdate := time.Now()

	for !done {
		ev := eventSource.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventInterrupt:
			switch ev.Data().(type) {
			case periodicUpdateInterruptCustomEvent:
				if repeatingMouseEvent != nil {
					view.HandleButton1PressEvent(repeatingMouseEvent.x, repeatingMouseEvent.y)
				}
				if time.Since(lastModelUpdate) >= 500*time.Millisecond {
					lastModelUpdate = time.Now()
					tetrisModel.PeriodicUpdate()
					view.Draw()
				}
			}
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				if !runningInWASM {
					quit()
				}
			case tcell.KeyLeft:
				tetrisModel.MoveCurrentPieceLeft()
				view.Draw()
			case tcell.KeyRight:
				tetrisModel.MoveCurrentPieceRight()
				view.Draw()
			case tcell.KeyUp:
				tetrisModel.RotateCurrentPiece()
				view.Draw()
			case tcell.KeyDown:
				tetrisModel.MoveCurrentPieceDown()
				view.Draw()
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'q':
					if !runningInWASM {
						quit()
					}
				case 'r':
					tetrisModel.Restart()
					view.Draw()
				case ' ':
					tetrisModel.DropCurrentPiece()
					view.Draw()
				case 'v':
					view.ToggleShowVersion()
					view.Draw()
				}
			}
		case *tcell.EventMouse:
			buttonMask := ev.Buttons()
			if (buttonMask & tcell.Button1) != 0 {
				x, y := ev.Position()
				repeatingMouseEvent = &mouseEventInfo{
					x: x,
					y: y,
				}
				view.HandleButton1PressEvent(x, y)
			} else {
				repeatingMouseEvent = nil
			}
		case *tcell.EventResize:
			view.HandleResizeEvent()
		}
	}

}

type periodicUpdateInterruptCustomEvent struct{}
