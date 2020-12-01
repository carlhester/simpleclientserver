package comms

import (
	"bufio"
	"fmt"

	"github.com/crucialcarl/simpleclientserver/server/user"
)

type Communicator struct {
}

func (c Communicator) SendMsgTo(msg string, users ...*user.User) {
	if msg[len(msg)-1] != '\n' {
		msg = msg + string('\n')
	}
	for _, p := range users {
		p.Inbox <- msg
		writer := bufio.NewWriter(p.Conn)
		_, err := writer.WriteString(msg)
		if err != nil {
			p.Close(err.Error())
		}
	}
}

func (c Communicator) ReceiveIncomingMessages(p user.User) {
	for {
		scanner := bufio.NewScanner(p.Conn)
		for scanner.Scan() {
			txt := scanner.Text()
			msg := message{
				src: p,
				txt: txt,
			}
			fmt.Printf("%+v\n", msg)
			p.Outbox <- txt
		}
	}
}
