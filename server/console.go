package main

import (
	"bufio"
	"fmt"
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

func setupConsole(id int, conn serverConsole, pList *playerList) player {
	console := player{
		id:    id,
		conn:  conn,
		name:  "CONSOLE",
		pList: pList,
	}
	return console
}

func consoleInput(p player) {
	fmt.Fprintf(os.Stdout, "> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := fmt.Sprintf("[%d] %s: %s", p.id, p.name, scanner.Text())
		if msg[len(msg)-1] != '\n' {
			msg = msg + string('\n')
		}
		sendMsgTo(nil, msg, p.pList.players...)
		fmt.Fprintf(os.Stdout, "> ")
	}
}
