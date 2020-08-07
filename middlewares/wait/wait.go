package wait

import (
	"time"

	"github.com/mvrilo/cam"
)

type Wait time.Duration

func New(t time.Duration) Wait {
	return Wait(t)
}

func (t Wait) Handle(f *cam.Frame) {
	time.Sleep(time.Duration(t))
}
