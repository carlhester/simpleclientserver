package player

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type communicator interface {
	SendMsgTo(string, ...Player)
	ListenForMessages(Player)
	EchoMessages(Player, *PlayerList)
}

// ClientFrontEnd is how the server interacts with the external clients
type ClientFrontEnd interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close() error
}

// Player ...
type Player struct {
	id    int
	Conn  ClientFrontEnd
	name  string
	Msgs  chan string
	PList *PlayerList
	comm  communicator
}

func SetupNewPlayer(conn net.Conn, id int, PlayerList *PlayerList, comm communicator) {
	var p *Player
	var msgs = make(chan string)
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
	p.comm.SendMsgTo(fmt.Sprintf("You are Player %d", id), *p)
	go p.comm.ListenForMessages(*p)
	go p.comm.EchoMessages(*p, PlayerList)
}

func getPlayerName(p *Player) error {
	p.comm.SendMsgTo("Hello! What is your name? ", *p)
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
