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
	prefix := fmt.Sprintf("[%d] %s: ", p.GetId(), p.GetName())
	for {
		scanner := bufio.NewScanner(p.Conn)
		for scanner.Scan() {
			txt := scanner.Text()
			switch txt {
			case "who":
				p.Msgs <- fmt.Sprintf("PlayerList\n==============\n")
				for i, player := range p.PList {
					p.Msgs <- fmt.Sprintf("[%d] %s", i, player.Name)
				}
				p.Msgs <- fmt.Sprintf("==============\n")
			default:
				p.Msgs <- prefix + txt + string('\n')
			}
		}
	}
}

func (c Communicator) EchoMessages(player player.Player, playerList player.PlayerList) {
	for {
		txt, ok := <-player.Msgs
		if ok {
			for _, p := range playerList {
				//	if p.Id != player.Id {
				c.SendMsgTo(txt, *p)
				//	}
			}
		}
	}
}
