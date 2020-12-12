package cam

// Device interfaces with the device
type Device interface {
	NewFrame() Frame
	Read(data interface{})
	Open(device interface{}) error
	Close() error
}
