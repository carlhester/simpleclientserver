package server

import (
	"net"

	"github.com/crucialcarl/simpleclientserver/server/comms"
	"github.com/crucialcarl/simpleclientserver/server/game"
)

type Server struct {
}

func (s Server) Run() {
	g := game.Game{}

	// init id incrementer
	id := make(chan int)
	go incrementer(id)

	// init error channel to receive errors from goroutines
	errChan := make(chan game.ClientErr)
	go errHandler(errChan)

	// Accepted connections go into a channel to be set up
	newConns := make(chan *net.Conn)
	addr := &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123}
	go comms.Listen(addr, newConns)
	for {
		conn := <-newConns
		go game.SetupNewPlayer(*conn, &g, <-id, &g.PlayerList, errChan)
	}

}

// errHandler receives errors from goroutines
func errHandler(err <-chan game.ClientErr) {
	e := <-err
	e.P.PList.Remove(*e.P)
	e.P.Conn.Close()
}

func incrementer(id chan<- int) {
	i := 0
	for {
		id <- i
		i++
	}
}
