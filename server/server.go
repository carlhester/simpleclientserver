package server

import (
	"net"
	"os"

	"github.com/crucialcarl/simpleclientserver/server/comms"
	"github.com/crucialcarl/simpleclientserver/server/player"
)

type Server struct {
	PlayerList *player.PlayerList
}

func (s Server) Run() {
	id := make(chan int)
	go incrementer(id)
	s.PlayerList = &player.PlayerList{
		Players: []player.Player{player.Player{
			Id: <-id,
			Conn: player.ServerConsole{
				Writer: os.Stdout,
				Reader: os.Stdin,
			},
			Name:  "CONSOLE",
			PList: s.PlayerList,
		},
		},
	}

	// Accepted connections go into a channel to be set up
	newConns := make(chan *net.Conn)
	addr := &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123}
	go comms.Listen(addr, newConns)
	for {
		conn := <-newConns
		comm := comms.Communicator{}
		go player.SetupNewPlayer(*conn, <-id, s.PlayerList, comm)
	}
}

func incrementer(id chan<- int) {
	i := 0
	for {
		id <- i
		i++
	}
}
