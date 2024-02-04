package controller

import (
	"sync/atomic"
	"time"

	"github.com/aaronriekenberg/go-tetris/model"
	"github.com/aaronriekenberg/go-tetris/utils"
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

type updateDuration struct {
	currentDuration time.Duration
}

var atomicUpdateDuration atomic.Pointer[updateDuration]

func setAtomicUpdateDuration(currentDuration time.Duration) {
	atomicUpdateDuration.Store(&updateDuration{
		currentDuration: currentDuration,
	})
}

func loadAtomicUpdateDuration() time.Duration {
	return atomicUpdateDuration.Load().currentDuration
}

func runEventLoop(
	eventSource view.ScreenEventSource,
	view view.View,
	tetrisModel model.TetrisModel,
) {
	setAtomicUpdateDuration(tetrisModel.UpdateDuration())

	go func() {
		for {
			sleepDuration := loadAtomicUpdateDuration()
			time.Sleep(sleepDuration)

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
				setAtomicUpdateDuration(tetrisModel.UpdateDuration())

				view.Draw()
			}
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				if !utils.RunningInWASM {
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
					if !utils.RunningInWASM {
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
				}
			}
		case *tcell.EventMouse:
			buttonMask := ev.Buttons()
			if (buttonMask & tcell.Button1) != 0 {
				x, y := ev.Position()
				view.HandleButton1PressEvent(x, y, ev.When())
			}
		case *tcell.EventResize:
			view.HandleResizeEvent()
		}
	}

}

type periodicUpdateInterruptCustomEvent struct{}
