package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"github.com/mvrilo/cam"
	"github.com/mvrilo/cam/middlewares/classifierdetect"
	"github.com/mvrilo/cam/middlewares/window"
	"gocv.io/x/gocv"
)

func main() {
	cascadeFile := os.Args[1]

	cam.Use(classifierdetect.New("eye", cascadeFile))
	cam.Handle(func(f *cam.Frame) {
		facesData := f.Get("eye")
		if facesData == nil {
			return
		}

		faces, ok := facesData.([]image.Rectangle)
		if !ok {
			return
		}

		for _, r := range faces {
			blue := color.RGBA{0, 0, 255, 0}
			gocv.Rectangle(&f.Data, r, blue, 3)

			size := gocv.GetTextSize("eye", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&f.Data, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}
	})
	cam.Use(window.New("cam example"))
	log.Fatal(cam.ListenAndServe(0, nil))
}
