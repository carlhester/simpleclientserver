package main

import (
	"os"
)

// the serverConsole is a special kind of client
type serverConsole struct {
	writer *os.File
	reader *os.File
}

func (s serverConsole) Close() error {
	os.Exit(1)
	return nil
}

func (s serverConsole) Write(b []byte) (n int, err error) {
	return s.writer.Write(b)
}

func (s serverConsole) Read(b []byte) (n int, err error) {

	return 0, nil
}
