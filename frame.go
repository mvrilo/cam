package cam

// Frame interfaces with the device frame
type Frame interface {
	Storage
	Close()
	Data() interface{}
}
