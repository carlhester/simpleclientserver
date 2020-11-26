package player

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type comm interface {
	sendMsgTo(string, ...Player)
	listenForMessages(Player)
	echoMessages(Player, *PlayerList)
}

// ClientFrontEnd is how the server interacts with the external clients
type ClientFrontEnd interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close() error
}

// a Player is a client with some labels
type Player struct {
	id    int
	Conn  ClientFrontEnd
	name  string
	Msgs  chan string
	PList *PlayerList
	comm  comm
}

func SetupNewPlayer(conn net.Conn, id int, PlayerList *PlayerList, comm comm) {
	var p *Player
	var msgs = make(chan string)
	log.Printf("Client connected: %s...\n", conn.RemoteAddr())
	p = &Player{
		id:    id,
		Conn:  conn,
		Msgs:  msgs,
		PList: PlayerList,
		comm:  comm,
	}
	err := getPlayerName(p)
	if err != nil {
		p.Close(err.Error())
		return
	}
	PlayerList.Add(*p)
	p.comm.sendMsgTo(fmt.Sprintf("You are Player %d", id), *p)
	go p.comm.listenForMessages(*p)
	go p.comm.echoMessages(*p, PlayerList)
}

func getPlayerName(p *Player) error {
	p.comm.sendMsgTo("Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.Conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	p.name = name[:len(name)-1]
	return nil
}

func (p *Player) Close(msg string) {
	log.Printf("%s\n", msg)
	p.PList.Remove(*p)
	p.Conn.Close()
}

func (p Player) GetId() int {
	return p.id
}

func (p Player) GetName() string {
	return p.name
}

/*
func (p Player) GetMsgs() []string {
	var msgs []string
	for m := range msgs {
		msgs = append(msgs, m)
	}
	return msgs

}
*/
