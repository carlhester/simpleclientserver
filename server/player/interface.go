package player

type communicator interface {
	SendMsgTo(string, ...*Player)
	ListenForMessages(Player)
	EchoMessages(Player, PlayerList)
}

// clientFrontEnd is how the server interacts with the external clients
type clientFrontEnd interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close() error
}
