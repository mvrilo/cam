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

func (w *Window) Handle(f cam.Frame) {
	data := f.Data()
	switch mat := data.(type) {
	case gocv.Mat:
		w.Window.IMShow(mat)
		w.Window.WaitKey(1)
	default:
	}
}
