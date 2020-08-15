package cam

import (
	"errors"

	"gocv.io/x/gocv"
)

var (
	ERR_CANT_READ   = errors.New("can't read device")
	ERR_EMPTY_FRAME = errors.New("empty frame")
)

// Camera is a ReadCloser interface based on gocv.VideoCapture
type Camera interface {
	Read(*gocv.Mat) bool
	Close() error
}

type HandlerFunc func(*Frame)

func (h HandlerFunc) Handle(frame *Frame) {
	h(frame)
}

// Handler interface expects a Handle func that receivs a frame
type Handler interface {
	Handle(*Frame)
}

// Cam struct with the video, device address and a list of handlers
type Cam struct {
	video Camera

	Device   interface{}
	Handlers []Handler
}

// Handle appends to the list of handlers a func with a signature that matches the HandlerFunc
func (c *Cam) Handle(h func(*Frame)) {
	c.Handlers = append(c.Handlers, HandlerFunc(h))
}

// Use appends to the list of handlers one or more Handlers
func (c *Cam) Use(handlers ...Handler) {
	c.Handlers = append(c.Handlers, handlers...)
}

func (c *Cam) handleFrame(frame *Frame) error {
	if ok := c.video.Read(&frame.Data); !ok {
		return ERR_CANT_READ
	}

	if frame.Data.Empty() {
		return ERR_EMPTY_FRAME
	}

	handlers := c.Handlers
	if len(handlers) == 0 {
		handlers = defaultHandlers
	}

	for _, handler := range handlers {
		if handler != nil {
			handler.Handle(frame)
		}
	}

	return nil
}

func (c *Cam) Serve() error {
	frame := newFrame()
	defer frame.Data.Close()
	for {
		err := c.handleFrame(frame)
		if err != nil {
			return err
		}
		frame.Index++
	}
}

func (c *Cam) ListenAndServe() error {
	video, err := gocv.OpenVideoCapture(c.Device)
	if err != nil {
		return err
	}
	c.video = video
	return c.Serve()
}

func (c *Cam) Close() error {
	if c.video != nil {
		return c.video.Close()
	}
	return nil
}

var defaultHandlers []Handler

func Use(h ...Handler) {
	defaultHandlers = append(defaultHandlers, h...)
}

func Handle(f func(f *Frame)) {
	defaultHandlers = append(defaultHandlers, HandlerFunc(f))
}

func ListenAndServe(device interface{}, handler Handler) error {
	c := &Cam{Device: device}
	if handler != nil {
		c.Handlers = []Handler{handler}
	}
	return c.ListenAndServe()
}
