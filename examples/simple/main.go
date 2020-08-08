package main

import (
	"image"
	"image/color"
	"log"

	"github.com/mvrilo/cam"
	"github.com/mvrilo/cam/middlewares/window"
	"gocv.io/x/gocv"
)

func main() {
	cam.Handle(func(f *cam.Frame) {
		text := "hello world"
		blue := color.RGBA{0, 0, 255, 0}
		gocv.PutText(&f.Data, text, image.Pt(200, 200), gocv.FontHersheyPlain, 10, blue, 8)
	})
	cam.Use(window.New("cam example"))
	log.Fatal(cam.ListenAndServe(0, nil))
}
