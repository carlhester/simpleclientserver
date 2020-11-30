package comms

import (
	"bufio"

	"github.com/crucialcarl/simpleclientserver/server/player"
)

type Communicator struct {
}

func (c Communicator) SendMsgTo(msg string, players ...*player.Player) {
	if msg[len(msg)-1] != '\n' {
		msg = msg + string('\n')
	}
	for _, p := range players {
		p.MsgsIn <- msg
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
	//prefix := fmt.Sprintf("[%d] %s: ", p.Id, p.Name)
	for {
		scanner := bufio.NewScanner(p.Conn)
		for scanner.Scan() {
			txt := scanner.Text()
			p.MsgsOut <- txt
		}
	}
}

func (c Communicator) EchoMessages(player player.Player, playerList player.PlayerList) {
	for {
		txt, ok := <-player.MsgsOut
		if ok {
			for _, p := range playerList {
				c.SendMsgTo(txt, p)
			}
		}
	}
}
