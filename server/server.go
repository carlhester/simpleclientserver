package server

import (
	"net"
	"os"

	"github.com/crucialcarl/simpleclientserver/server/comms"
	"github.com/crucialcarl/simpleclientserver/server/user"
)

type Server struct {
	UserList user.UserList
}

func (s Server) Run() {
	id := make(chan int)
	go incrementer(id)

	s.UserList = user.NewUserList()

	console := &user.User{
		Id: <-id,
		Conn: user.ServerConsole{
			Writer: os.Stdout,
			Reader: os.Stdin,
		},
		Name:  "CONSOLE",
		PList: s.UserList,
	}

	s.UserList[console.Id] = console

	newConns := make(chan *net.Conn)
	addr := &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123}

	go comms.Listen(addr, newConns)
	for {
		conn := <-newConns
		comm := comms.Communicator{}
		go user.SetupNewUser(*conn, <-id, s.UserList, comm)
	}
}

func incrementer(id chan<- int) {
	i := 0
	for {
		id <- i
		i++
	}
}
