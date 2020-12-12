package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type simpleServer struct {
	listener
	userlist *userlist
	msgsChan chan message // channel of messages inbound from clients
}

func (s simpleServer) handleMsgs() {
	for {
		msg := <-s.msgsChan
		msg.txt = strings.TrimSpace(msg.txt)
		fmt.Printf("From %s: %+s (%q) (%+v)\n", msg.src.name, msg.txt, msg.txt, msg)
		if strings.HasPrefix(msg.txt, "/") {
			s.handleCommand(msg)
		}

		_, err := fmt.Fprintf(msg.src, "> ")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s simpleServer) handleCommand(msg message) {
	switch msg.txt {
	case "/who":
		result := fmt.Sprintf("NAME\tID\tTIME\n")
		for _, u := range s.userlist.users {
			result = result + fmt.Sprintf("%s\t%d\t%s\n", u.name, u.id, u.loginTime.Format("Mon Jan 2 15:04:05 MST 2006"))
			log.Print(result)
			_, err := fmt.Fprintf(msg.src, result)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

func newSimpleServer(c config) *simpleServer {
	addr := &net.TCPAddr{
		IP:   net.ParseIP(c.ip),
		Port: c.port,
	}

	return &simpleServer{
		userlist: &userlist{},
		listener: listener{
			addr:     addr,
			newConns: make(chan *net.Conn),
		},
		msgsChan: make(chan message),
	}
}

func (s *simpleServer) run() error {
	id := int(0)
	go s.listen()
	go s.handleMsgs()
	for {
		conn := <-s.newConns
		u := newUser(id, conn, s.msgsChan)
		s.addToUserList(u)
		go u.process()
		id++
	}
	return nil
}

func (s *simpleServer) listen() {
	log.Printf("listening on %s", s.addr)
	listener, err := net.ListenTCP("tcp", s.addr)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Client connected: %s...\n", conn.RemoteAddr())
		s.newConns <- &conn
	}
}

func (s *simpleServer) addToUserList(u *user) {
	s.userlist.lock.Lock()
	s.userlist.users = append(s.userlist.users, u)
	s.userlist.lock.Unlock()
}

func (s *simpleServer) removeFromUserList(u *user) {
	results := []*user{}
	for _, each := range s.userlist.users {
		if each != u {
			results = append(results, each)
		}
	}
	s.userlist.lock.Lock()
	s.userlist.users = results
	s.userlist.lock.Unlock()
}
