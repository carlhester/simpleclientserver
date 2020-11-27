package console

import (
	"os"
)

// the ServerConsole is a special kind of client
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

/*
func consoleInput(p game.player) {
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
*/
