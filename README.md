# cam

`cam` is a high level package built on top of [gocv](https://gocv.io/) for working with the camera and its frames.
It has an api familiar to `net/http` and brings a set of middlewares for frame handling.

## Installation

```
go get github.com/mvrilo/cam
```

## Example

```
package main

import (
	"image"
	"image/color"
	"log"

	"github.com/mvrilo/cam"
	"github.com/mvrilo/cam/middlewares/recorder"
	"gocv.io/x/gocv"
)

func main() {
	cam.Handle(func(f *cam.Frame) {
		text := "hello world"
		blue := color.RGBA{0, 0, 255, 0}
		gocv.PutText(&f.Data, text, image.Pt(200, 200), gocv.FontHersheyPlain, 10, blue, 8)
	})
	cam.Use(recorder.New("./out.avi"))
	log.Fatal(cam.ListenAndServe(0, nil))
}
```

## Author

Murilo Santana <<mvrilo@gmail.com>>

## License

MIT
