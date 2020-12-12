package gocv

import (
	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

// frame is a gocv frame
type frame struct {
	cam.Storage
	data  gocv.Mat
	index int
}

func newFrame() *frame {
	return &frame{
		Storage: cam.NewStore(),
		data:    gocv.NewMat(),
	}
}

func (f *frame) Close() {
	f.data.Close()
}

func (f *frame) Data() interface{} {
	return f.data
}
