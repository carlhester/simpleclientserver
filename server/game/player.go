package game

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// a player is a client with some labels
type player struct {
	id    int
	Conn  ClientFrontEnd
	name  string
	msgs  chan string
	PList *PlayerList
}

func SetupNewPlayer(conn net.Conn, game *Game, id int, PlayerList *PlayerList, errChan chan<- ClientErr) {
	var newPlayer *player
	var msgs = make(chan string)
	log.Printf("Client connected: %s...\n", conn.RemoteAddr())
	newPlayer = &player{
		id:    id,
		Conn:  conn,
		msgs:  msgs,
		PList: PlayerList,
	}
	err := getPlayerName(newPlayer)
	if err != nil {
		clErr := ClientErr{
			P:   newPlayer,
			err: err,
		}
		errChan <- clErr
		return
	}
	game.PlayerList.Add(*newPlayer)
	sendMsgTo(nil, fmt.Sprintf("You are player %d", id), *newPlayer)
	go listenForMessages(*newPlayer)
	go echoMessages(errChan, *newPlayer, &game.PlayerList)
}

func getPlayerName(p *player) error {
	sendMsgTo(nil, "Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.Conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	p.name = name[:len(name)-1]
	return nil
}
