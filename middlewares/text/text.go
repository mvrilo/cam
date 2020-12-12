package text

import (
	"image"
	"image/color"

	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

type Text struct {
	Value     string
	Position  image.Point
	Color     color.RGBA
	FontName  gocv.HersheyFont
	FontScale float64
	Thickness int
	LineType  gocv.LineType
	Bottom    bool

	OnFrameSize func(t *Text, s []int)
}

func defaults(t *Text) {
	if t.FontScale == 0.0 {
		t.FontScale = 6.0
	}

	if t.Thickness == 0 {
		t.Thickness = 12
	}

	if t.LineType == 0 {
		t.LineType = 16
	}
}

func (t *Text) Handle(f cam.Frame) {
	defaults(t)

	frameData := f.Data()
	switch data := frameData.(type) {
	case gocv.Mat:
		t.OnFrameSize(t, data.Size())
		gocv.PutTextWithParams(
			&data,
			t.Value,
			t.Position,
			t.FontName,
			t.FontScale,
			t.Color,
			t.Thickness,
			t.LineType,
			t.Bottom,
		)
	default:
	}
}
