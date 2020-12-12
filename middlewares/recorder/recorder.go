package recorder

import (
	"log"

	"github.com/mvrilo/cam"
	"github.com/pkg/errors"
	"gocv.io/x/gocv"
)

type Recorder struct {
	vw     *gocv.VideoWriter
	loaded bool

	Output string
	Codec  string
	FPS    float32
}

func New(output string) *Recorder {
	return &Recorder{
		Output: output,
		Codec:  "MJPG",
		FPS:    10.0,
	}
}

func (r *Recorder) handleGocv(data gocv.Mat) {
	if !r.loaded {
		vw, err := gocv.VideoWriterFile(
			r.Output,
			r.Codec,
			float64(r.FPS),
			data.Cols(),
			data.Rows(),
			true,
		)
		if err != nil {
			log.Println(errors.Wrap(err, "recorder"))
			return
		}
		r.vw = vw
		r.loaded = true
	}

	go func() {
		if r.vw == nil || !r.vw.IsOpened() {
			log.Println(errors.New("recorder: video writer not opened"))
			return
		}

		err := r.vw.Write(data)
		if err != nil {
			log.Println(errors.Wrap(err, "recorder"))
		}
	}()
}

func (r *Recorder) Handle(f cam.Frame) {
	frameData := f.Data()
	switch data := frameData.(type) {
	case gocv.Mat:
		r.handleGocv(data)
	default:
	}
}
