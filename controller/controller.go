package controller

import (
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

func runEventLoop(
	eventSource view.ScreenEventSource,
	view view.View,
	tetrisModel model.TetrisModel,
) {
	periodicUpdateTicker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			<-periodicUpdateTicker.C

			eventSource.PostEvent(tcell.NewEventInterrupt(periodicUpdateInterruptCustomEvent{}))
		}
	}()

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

	for !done {
		ev := eventSource.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventInterrupt:
			switch ev.Data().(type) {
			case periodicUpdateInterruptCustomEvent:
				tetrisModel.PeriodicUpdate()
				view.Draw()
			}
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				quit()
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
					quit()
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
		case *tcell.EventResize:
			view.HandleResizeEvent()
		}
	}

}

type periodicUpdateInterruptCustomEvent struct{}
