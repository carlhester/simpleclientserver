package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// game handles the high level "global" state of the game
type game struct {
	playerList
}

type clientErr struct {
	p   *player
	err error
}

func (g game) Run() {
	// initialize id incrementer
	id := make(chan int)
	go incrementer(id)

	// initialize error channel to receive errors from goroutines
	errChan := make(chan clientErr)
	go errHandler(errChan)

	// the serverConsole uses standard in/out
	serverConsole := serverConsole{
		writer: os.Stdin,
		reader: os.Stdout,
	}

	// players keeps a list of active players
	// the initial player is the serverConsole
	consoleId := <-id
	console := player{
		id:    consoleId,
		conn:  serverConsole,
		name:  "CONSOLE",
		pList: &g.playerList,
	}
	g.playerList.add(console)
	go consoleInput(console)

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
		go setupNewPlayer(conn, &g, id, &g.playerList, errChan)
	}
}

func consoleInput(p player) {
	fmt.Fprintf(os.Stdout, "> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := p.name + ": " + scanner.Text()
		if msg[len(msg)-1] != '\n' {
			msg = msg + string('\n')
		}
		sendMsgTo(nil, msg, p.pList.players...)
		fmt.Fprintf(os.Stdout, "> ")
	}
}

// errHandler receives errors from goroutines
func errHandler(err <-chan clientErr) {
	e := <-err
	e.p.pList.remove(*e.p)
	e.p.conn.Close()
}

func incrementer(id chan<- int) {
	i := 0
	for {
		id <- i
		i++
	}
}
