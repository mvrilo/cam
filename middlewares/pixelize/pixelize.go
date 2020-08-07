package pixelize

import (
	"image"

	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

type Pixelize struct {
	size int
}

func New(size int) *Pixelize {
	return &Pixelize{size}
}

func (p *Pixelize) Handle(f *cam.Frame) {
	if p.size == 0 {
		return
	}

	frameSize := f.Data.Size()
	tmp := gocv.NewMat()

	gocv.Resize(f.Data, &tmp, image.Point{X: p.size, Y: p.size}, 0.0, 0.0, gocv.InterpolationLinear)
	gocv.Resize(tmp, &f.Data, image.Point{X: frameSize[1], Y: frameSize[0]}, 0.0, 0.0, gocv.InterpolationNearestNeighbor)
}
