package user

import (
	"os"
)

// ServerConsole is a special kind of client that satisfies the clientFrontEnd interface
type ServerConsole struct {
	Writer *os.File
	Reader *os.File
}

func (s ServerConsole) Close() error {
	os.Exit(1)
	return nil
}

func (s ServerConsole) Write(b []byte) (n int, err error) {
	return s.Writer.Write(b)
}

func (s ServerConsole) Read(b []byte) (n int, err error) {

	return 0, nil
}
