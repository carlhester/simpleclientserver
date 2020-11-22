package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// a player is a client with some labels
type player struct {
	id    int
	conn  clientFrontEnd
	name  string
	msgs  chan string
	pList *playerList
}

func setupNewPlayer(conn net.Conn, game *game, id int, playerList *playerList) {
	var newPlayer *player
	var msgs = make(chan string)
	log.Printf("Client connected: %s...\n", conn.RemoteAddr())
	newPlayer = &player{
		id:    id,
		conn:  conn,
		msgs:  msgs,
		pList: playerList,
	}
	getPlayerName(newPlayer)
	game.playerList.add(*newPlayer)
	sendMsgTo(fmt.Sprintf("You are player %d", id), *newPlayer)
	go listenForMessages(*newPlayer)
	go echoMessages(*newPlayer, &game.playerList)
}

func getPlayerName(p *player) {
	sendMsgTo("Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.conn)
	name, _ := reader.ReadString('\n')
	p.name = name[:len(name)-1]
	log.Println(name)
}
