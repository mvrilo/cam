package classifierdetect

import (
	"image"

	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

type ClassifierDetect struct {
	Key        string
	Classifier gocv.CascadeClassifier
}

func New(storeKey, file string) *ClassifierDetect {
	c := gocv.NewCascadeClassifier()
	if !c.Load(file) {
		panic("couldn't load cascade file")
	}

	return &ClassifierDetect{
		Classifier: c,
		Key:        storeKey,
	}
}

func (fd *ClassifierDetect) Handle(f cam.Frame) {
	var rects []image.Rectangle
	frameData := f.Data()
	switch data := frameData.(type) {
	case gocv.Mat:
		rects = fd.Classifier.DetectMultiScale(data)
	}
	f.Set(fd.Key, rects)
}
