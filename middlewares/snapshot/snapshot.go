package snapshot

import (
	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

type Snapshot struct {
	Output string
}

func New(output string) *Snapshot {
	return &Snapshot{output}
}

func (s *Snapshot) Handle(f cam.Frame) {
	frameData := f.Data()
	switch data := frameData.(type) {
	case gocv.Mat:
		go gocv.IMWrite(s.Output, data)
	default:
	}
}
