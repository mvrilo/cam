package main

import (
	"image"
	"image/color"
	"log"

	"github.com/mvrilo/cam"
	"github.com/mvrilo/cam/middlewares/window"
	"gocv.io/x/gocv"

	_ "github.com/mvrilo/cam/gocv"
)

func main() {
	cam.Handle(func(f cam.Frame) {
		text := "hello world"
		blue := color.RGBA{0, 0, 255, 0}
		data := f.Data()
		if mat, ok := data.(gocv.Mat); ok {
			gocv.PutText(&mat, text, image.Pt(200, 200), gocv.FontHersheyPlain, 10, blue, 8)
		}
	})
	cam.Use(window.New("cam example"))
	log.Fatal(cam.ListenAndServe(0, nil))
}
