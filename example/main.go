package main

import (
	"image"
	"image/color"
	"log"
	"net/http"
	"time"

	"github.com/mvrilo/cam"
	"github.com/mvrilo/cam/middlewares/pixelize"
	"github.com/mvrilo/cam/middlewares/recorder"
	"github.com/mvrilo/cam/middlewares/snapshot"
	"github.com/mvrilo/cam/middlewares/streamer"
	"github.com/mvrilo/cam/middlewares/wait"
	"github.com/mvrilo/cam/middlewares/window"
	"gocv.io/x/gocv"
)

func main() {
	cam.Use(pixelize.New(64))

	cam.Handle(func(f *cam.Frame) {
		text := "hello world"
		blue := color.RGBA{0, 0, 255, 0}
		gocv.PutText(&f.Data, text, image.Pt(200, 200), gocv.FontHersheyPlain, 10, blue, 8)
	})

	streaming := streamer.New("/cam.mjpg")
	cam.Use(
		streaming,
		window.New("cam example"),
		snapshot.New("./out.jpg"),
		recorder.New("./out.avi"),
		wait.New(100*time.Millisecond),
	)

	go func() {
		println("streaming at http://localhost:8000/cam.mjpg")
		log.Fatal(http.ListenAndServe(":8000", streaming))
	}()

	log.Fatal(cam.ListenAndServe(0, nil))
}
