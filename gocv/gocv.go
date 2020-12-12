package gocv

import (
	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

func init() {
	cam.Register(&Gocv{})
}

type Gocv struct {
	video *gocv.VideoCapture
}

func (g *Gocv) Open(device interface{}) error {
	video, err := gocv.OpenVideoCapture(device)
	if err != nil {
		return err
	}
	g.video = video
	return nil
}

func (g *Gocv) Close() error {
	return g.video.Close()
}

func (g *Gocv) Read(data interface{}) {
	mat, ok := data.(gocv.Mat)
	if ok {
		g.video.Read(&mat)
	}
}

func (g *Gocv) NewFrame() cam.Frame {
	return newFrame()
}
