package cam

type HandlerFunc func(Frame)

func (h HandlerFunc) Handle(frame Frame) {
	h(frame)
}

// Handler interface expects a Handle func that receivs a frame
type Handler interface {
	Handle(Frame)
}
