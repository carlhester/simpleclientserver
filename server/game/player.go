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
	conn  ClientFrontEnd
	name  string
	msgs  chan string
	pList *playerList
}

func setupNewPlayer(conn net.Conn, game *Game, id int, playerList *playerList, errChan chan<- clientErr) {
	var newPlayer *player
	var msgs = make(chan string)
	log.Printf("Client connected: %s...\n", conn.RemoteAddr())
	newPlayer = &player{
		id:    id,
		conn:  conn,
		msgs:  msgs,
		pList: playerList,
	}
	err := getPlayerName(newPlayer)
	if err != nil {
		clErr := clientErr{
			p:   newPlayer,
			err: err,
		}
		errChan <- clErr
		return
	}
	game.playerList.Add(*newPlayer)
	sendMsgTo(nil, fmt.Sprintf("You are player %d", id), *newPlayer)
	go listenForMessages(*newPlayer)
	go echoMessages(errChan, *newPlayer, &game.playerList)
}

func getPlayerName(p *player) error {
	sendMsgTo(nil, "Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	p.name = name[:len(name)-1]
	return nil
}
