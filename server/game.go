package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

// game handles the high level "global" state of the game
type game struct {
	playerList
}

func (g game) Run() {
	// initialize id incrementer
	id := make(chan int)
	go incrementer(id)

	// the serverConsole uses standard in/out
	serverConsole := serverConsole{
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	}

	// players keeps a list of active players
	// the initial player is the serverConsole
	consoleId := <-id
	console := player{
		id:    consoleId,
		conn:  serverConsole,
		name:  "server",
		pList: &g.playerList,
	}
	g.playerList.add(console)

	// create players by listening for network connects
	addr := &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123}
	log.Printf("listening on %s", addr)
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	// accept network connections and assign players
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		id := <-id
		go setupNewPlayer(conn, &g, id, &g.playerList)
	}

}
