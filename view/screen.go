package view

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v3"
	_ "github.com/gdamore/tcell/v3/encoding"
)

type ScreenEventSource interface {
	EventQ() chan tcell.Event
	Fini()
}

var _ ScreenEventSource = (tcell.Screen)(nil)

func NewScreen() tcell.Screen {
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

	screen.EnableMouse()

	return screen
}
