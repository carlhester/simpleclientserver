package comms

import (
	"bufio"
	"fmt"

	"github.com/crucialcarl/simpleclientserver/server/player"
)

type Communicator struct {
}

func (c Communicator) SendMsgTo(msg string, players ...player.Player) {
	if msg[len(msg)-1] != '\n' {
		msg = msg + string('\n')
	}
	for _, p := range players {
		writer := bufio.NewWriter(p.Conn)
		_, err := writer.WriteString(msg)
		if err != nil {
			p.Close(err.Error())
		}
		err = writer.Flush()
		if err != nil {
			p.Close(err.Error())
		}
	}
}

func (c Communicator) ListenForMessages(p player.Player) {
	for {
		scanner := bufio.NewScanner(p.Conn)
		for scanner.Scan() {
			prefix := fmt.Sprintf("[%d] %s: ", p.GetId(), p.GetName())
			p.Msgs <- prefix + scanner.Text() + "\n"
		}
	}
}

func (c Communicator) EchoMessages(player player.Player, playerList *player.PlayerList) {
	for {
		txt, ok := <-player.Msgs
		if ok {
			for _, p := range playerList.Get() {
				if p != player {
					c.SendMsgTo(txt, p)
				}
			}
		}
	}
}
