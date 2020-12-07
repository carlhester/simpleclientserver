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
	msgsChan  chan message // channel of messages inbound from clients
}

func (s simpleServer) handleMsgs() {
	for {
		msg := <-s.msgsChan
		fmt.Printf("From %s: %+s (%+v)\n", msg.src.name, msg.txt, msg)
		if strings.HasPrefix(msg.txt, "/") {
			s.handleCommand(msg)
		}
	}
}

func (s simpleServer) handleCommand(msg message) {
	switch msg.txt {
	case "/who":
		result := fmt.Sprintf("%+v\n", s.userlist)
		log.Print(result)
		_, err := fmt.Fprintf(msg.src, result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func newSimpleServer(c config) *simpleServer {
	userlist := &userlist{}
	newConns := make(chan *net.Conn)
	msgsChan := make(chan message)

	addr := &net.TCPAddr{
		IP:   net.ParseIP(c.ip),
		Port: c.port,
	}

	return &simpleServer{
		userlist: userlist,
		listener: listener{
			addr:     addr,
			newConns: newConns,
		},
		msgsChan: msgsChan,
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
