package game

// ClientFrontEnd is how the server interacts with the external clients
type ClientFrontEnd interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close() error
}
