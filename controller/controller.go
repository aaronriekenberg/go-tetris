package controller

import (
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

func runEventLoop(
	eventSource view.ScreenEventSource,
	view view.View,
	tetrisModel model.TetrisModel,
) {
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

	eventChannel := make(chan tcell.Event)

	go eventSource.ChannelEvents(eventChannel, make(chan struct{}))

	periodicUpdateChannel := time.After(tetrisModel.PeriodicUpdateDuration())

	for !done {
		select {
		case <-periodicUpdateChannel:
			tetrisModel.PeriodicUpdate()
			view.Draw()

			periodicUpdateChannel = time.After(tetrisModel.PeriodicUpdateDuration())
		case ev := <-eventChannel:
			switch ev := ev.(type) {
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
}
