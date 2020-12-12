package cam

// Cam struct with the video, device address and a list of handlers
type Cam struct {
	Addr     interface{}
	Device   Device
	Handlers []Handler
}

// Register sets a video device
func (c *Cam) Register(video Device) {
	c.Device = video
}

// Handle appends to the list of handlers a func with a signature that matches the HandlerFunc
func (c *Cam) Handle(h func(Frame)) {
	c.Handlers = append(c.Handlers, HandlerFunc(h))
}

// Use appends to the list of handlers one or more Handlers
func (c *Cam) Use(handlers ...Handler) {
	c.Handlers = append(c.Handlers, handlers...)
}

func (c *Cam) Serve() {
	frame := c.Device.NewFrame()
	defer frame.Close()
	for {
		data := frame.Data()
		c.Device.Read(data)

		handlers := c.Handlers
		if len(handlers) == 0 {
			handlers = defaultHandlers
		}

		for _, handler := range handlers {
			if handler != nil {
				handler.Handle(frame)
			}
		}
	}
}

func (c *Cam) ListenAndServe() error {
	err := c.Device.Open(c.Addr)
	if err != nil {
		return err
	}
	c.Serve()
	return nil
}

func (c *Cam) Close() error {
	if c.Device != nil {
		return c.Device.Close()
	}
	return nil
}

var defaultHandlers []Handler

var defaultDevice Device = nil

func Register(device Device) {
	defaultDevice = device
}

func Use(h ...Handler) {
	defaultHandlers = append(defaultHandlers, h...)
}

func Handle(f func(f Frame)) {
	defaultHandlers = append(defaultHandlers, HandlerFunc(f))
}

func ListenAndServe(addr interface{}, handler Handler) error {
	c := &Cam{
		Addr:   addr,
		Device: defaultDevice,
	}
	if handler != nil {
		c.Handlers = []Handler{handler}
	}
	return c.ListenAndServe()
}
