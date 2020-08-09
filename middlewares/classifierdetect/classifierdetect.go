package classifierdetect

import (
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

func (fd *ClassifierDetect) Handle(f *cam.Frame) {
	rects := fd.Classifier.DetectMultiScale(f.Data)
	f.Store.Set(fd.Key, rects)
}
