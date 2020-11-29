package player

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Player ...
type Player struct {
	Id       int
	Conn     clientFrontEnd
	Name     string
	MsgsIn   chan string // Msgs destined for this Player
	MsgsOut  chan string // Msgs from this player
	PList    PlayerList
	comm     communicator
	location int
}

func (p *Player) Close(msg string) {
	log.Printf("%s\n", msg)
	delete(p.PList, p.Id)
	p.Conn.Close()
}

func SetupNewPlayer(conn net.Conn, id int, PlayerList PlayerList, comm communicator) {
	var p *Player
	var msgs = make(chan string)
	p = &Player{
		Id:     id,
		Conn:   conn,
		MsgsIn: msgs,
		PList:  PlayerList,
		comm:   comm,
	}
	err := getPlayerName(p)
	if err != nil {
		p.Close(err.Error())
		return
	}
	PlayerList[p.Id] = p
	p.comm.SendMsgTo(fmt.Sprintf("You are Player %d", id), *p)
	go p.comm.ListenForMessages(*p)
	go p.receiveMsgs()
}

func getPlayerName(p *Player) error {
	p.comm.SendMsgTo("Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.Conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	p.Name = name[:len(name)-1]
	return nil
}

func (p *Player) receiveMsgs() {
	for {
		txt, ok := <-p.MsgsIn
		if ok {
			fmt.Println(txt)
			writer := bufio.NewWriter(p.Conn)
			_, err := writer.WriteString(txt)
			if err != nil {
				p.Close(err.Error())
			}
			err = writer.Flush()
			if err != nil {
				p.Close(err.Error())
			}

		}
	}
}
