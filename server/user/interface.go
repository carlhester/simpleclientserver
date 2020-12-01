package user

type communicator interface {
	SendMsgTo(string, ...*User)
	ReceiveIncomingMessages(User)
}

// clientFrontEnd is how the server interacts with the external clients
type clientFrontEnd interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close() error
}
