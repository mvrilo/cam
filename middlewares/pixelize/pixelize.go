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

func (p *Pixelize) Handle(f cam.Frame) {
	if p.size == 0 {
		return
	}

	frameData := f.Data()
	switch data := frameData.(type) {
	case gocv.Mat:
		frameSize := data.Size()
		tmp := gocv.NewMat()
		gocv.Resize(data, &tmp, image.Point{X: p.size, Y: p.size}, 0.0, 0.0, gocv.InterpolationLinear)
		gocv.Resize(tmp, &data, image.Point{X: frameSize[1], Y: frameSize[0]}, 0.0, 0.0, gocv.InterpolationNearestNeighbor)
	default:
	}
}
