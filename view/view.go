package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type AppView struct {
	View tview.Primitive
}

type VimMotionCapture struct {
	JCapture func()
	KCapture func()
	HCapture func()
	LCapture func()
}

func (motionCapture VimMotionCapture) InputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			return handleIfNotNull(motionCapture.JCapture, event)
		case 'k':
			return handleIfNotNull(motionCapture.KCapture, event)
		case 'h':
			return handleIfNotNull(motionCapture.HCapture, event)
		case 'l':
			return handleIfNotNull(motionCapture.LCapture, event)
		}
		return event
	}
}

func handleIfNotNull(handler func(), event *tcell.EventKey) *tcell.EventKey {
	if handler != nil {
		handler()
		return nil
	}
	return event
}
