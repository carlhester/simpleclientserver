package main

import (
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

	// the initial player is the serverConsole
	console := setupConsole(<-id, serverConsole, &g.playerList)
	go consoleInput(console)
	g.playerList.add(console)

	newConns := make(chan *net.Conn)

	// create players by listening for network connects
	listener := listener{
		addr: &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123},
	}
	go listener.Listen(newConns)
	for {
		conn := <-newConns
		go setupNewPlayer(*conn, &g, <-id, &g.playerList, errChan)
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
