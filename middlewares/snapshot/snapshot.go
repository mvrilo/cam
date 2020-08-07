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

func (s *Snapshot) Handle(f *cam.Frame) {
	go gocv.IMWrite(s.Output, f.Data)
}
