package view

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"
)

type Screen interface {
	tcellScreen() tcell.Screen
	PollEvent() tcell.Event
	PostEvent(ev tcell.Event) error
	Fini()
}

type screen struct {
	tcell.Screen
}

func (s *screen) tcellScreen() tcell.Screen {
	return s
}

func NewScreen() Screen {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	tcellScreen, e := tcell.NewScreen()

	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	if e = tcellScreen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	return &screen{
		Screen: tcellScreen,
	}
}
