package ocr

import (
	"github.com/mvrilo/cam"
	"github.com/otiai10/gosseract"
	"gocv.io/x/gocv"
)

type OCR struct {
	cli *gosseract.Client
}

func New() *OCR {
	return &OCR{
		cli: gosseract.NewClient(),
	}
}

func (w *OCR) Handle(f cam.Frame) {
	frameData := f.Data()
	switch data := frameData.(type) {
	case gocv.Mat:
		w.cli.SetImageFromBytes((&data).ToBytes())
		text, _ := w.cli.Text()
		if text != "" {
			println(text)
		}
	default:
	}
}
