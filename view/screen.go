package view

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"
)

type ScreenEventSource interface {
	PollEvent() tcell.Event
	PostEvent(ev tcell.Event) error
	Fini()
}

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

	return screen
}
