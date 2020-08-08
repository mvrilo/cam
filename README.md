# cam

[![GoDoc](https://godoc.org/github.com/mvrilo/cam?status.svg)](https://godoc.org/github.com/mvrilo/cam)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvrilo/cam)](https://goreportcard.com/report/github.com/mvrilo/cam)

`cam` is a package built on top of [gocv](https://gocv.io/) that provides a high level api for working with the camera frames.
It has a similar api to `net/http` and `gliderlabs/ssh` plus a set of builtin middlewares for easy composition.

## Dependencies

OpenCV is the only dependency. You can see instructions for installation on any platform on [the gocv site](https://gocv.io/getting-started/).

## Installation

```
go get github.com/mvrilo/cam
```

## Example

```go
package main

import (
	"image"
	"image/color"
	"log"

	"github.com/mvrilo/cam"
	"github.com/mvrilo/cam/middlewares/recorder"
	"github.com/mvrilo/cam/middlewares/window"
	"gocv.io/x/gocv"
)

func main() {
	cam.Handle(func(f *cam.Frame) {
		text := "hello world"
		blue := color.RGBA{0, 0, 255, 0}
		gocv.PutText(&f.Data, text, image.Pt(200, 200), gocv.FontHersheyPlain, 10, blue, 8)
	})
	cam.Use(
		window.New("Hello world, Cam!"),
		recorder.New("./out.avi"),
	)
	log.Fatal(cam.ListenAndServe(0, nil))
}
```

## Author

Murilo Santana <<mvrilo@gmail.com>>

## License

MIT
