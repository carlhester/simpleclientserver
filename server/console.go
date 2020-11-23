package main

import (
	"bufio"
	"os"
)

// the serverConsole is a special kind of client
type serverConsole struct {
	*bufio.Reader
	*bufio.Writer
}

func (s serverConsole) Close() error {
	os.Exit(1)
	return nil
}
