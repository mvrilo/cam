package streamer

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/mvrilo/cam"
	"gocv.io/x/gocv"
)

var Boundary = "CAM"

const payload = "\r\n--%s\r\nContent-Type: image/jpeg\r\nContent-length: %d\r\n\r\n%s\r\n\r\n"

type Streamer struct {
	path  string
	size  int
	frame []byte
	mu    sync.Mutex
}

func New(path string) *Streamer {
	return &Streamer{
		path: path,
	}
}

func (s *Streamer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != s.path {
		return
	}

	w.Header().Set("Transfer-Encoding", "deflate")
	w.Header().Set("Content-Type", "multipart/x-mixed-replace;boundary="+Boundary)

	for {
		s.mu.Lock()
		fmt.Fprintf(w, payload, Boundary, len(s.frame), s.frame)
		s.mu.Unlock()
	}
}

func (s *Streamer) Handle(f cam.Frame) {
	s.mu.Lock()
	defer s.mu.Unlock()

	frameData := f.Data()
	switch data := frameData.(type) {
	case gocv.Mat:
		body, err := gocv.IMEncodeWithParams(gocv.JPEGFileExt, data, []int{gocv.IMWriteJpegQuality, 70})
		if err != nil {
			log.Println(err)
			return
		}
		s.frame = body
	default:
	}
}
