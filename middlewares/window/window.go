package window

import (
	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

type Window struct {
	*gocv.Window
}

func New(name string) *Window {
	return &Window{
		gocv.NewWindow(name),
	}
}

func (w *Window) Handle(f *cam.Frame) {
	w.Window.IMShow(f.Data)
	w.Window.WaitKey(1)
}
