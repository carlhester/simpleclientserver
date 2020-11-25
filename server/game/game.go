package game

import (
	"net"

	"github.com/crucialcarl/simpleclientserver/server/comms"
)

// game handles the high level "global" state of the game
type Game struct {
	playerList
}

func (g Game) Run() {
	// initialize id incrementer
	id := make(chan int)
	go incrementer(id)

	// initialize error channel to receive errors from goroutines
	errChan := make(chan clientErr)
	go errHandler(errChan)

	// Accepted connections go into a channel to be set up
	newConns := make(chan *net.Conn)
	addr := &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123}
	go comms.Listen(addr, newConns)
	for {
		conn := <-newConns
		go setupNewPlayer(*conn, &g, <-id, &g.playerList, errChan)
	}
}

// errHandler receives errors from goroutines
func errHandler(err <-chan clientErr) {
	e := <-err
	e.p.pList.Remove(*e.p)
	e.p.conn.Close()
}

func incrementer(id chan<- int) {
	i := 0
	for {
		id <- i
		i++
	}
}
