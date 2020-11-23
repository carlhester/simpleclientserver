package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/crucialcarl/simpleclientserver/server/frontend"
)

// a player is a client with some labels
type player struct {
	id    int
	conn  frontend.ClientFrontEnd
	name  string
	msgs  chan string
	pList *playerList
}

func setupNewPlayer(conn net.Conn, game *game, id int, playerList *playerList, errChan chan<- clientErr) {
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
	game.playerList.add(*newPlayer)
	sendMsgTo(fmt.Sprintf("You are player %d", id), *newPlayer)
	go listenForMessages(*newPlayer)
	go echoMessages(*newPlayer, &game.playerList)
}

func getPlayerName(p *player) error {
	sendMsgTo("Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	p.name = name[:len(name)-1]
	return nil
}
