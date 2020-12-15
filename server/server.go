package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type simpleServer struct {
	listener
	userlist  *userlist
	msgsChan  chan message // channel of messages inbound from clients
	commands  map[string]commandHandler
	startTime time.Time
	roomlist  []int
}

func (s simpleServer) roomExists(id int) bool {
	for _, i := range s.roomlist {
		if i == id {
			return true
		}
	}
	return false
}

func (s simpleServer) usersInRoom(id int) []*user {
	result := []*user{}
	for _, u := range s.userlist.users {
		if u.room == id {
			result = append(result, u)
		}
	}
	return result
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
	cmd := strings.Split(msg.txt, " ")
	handler, ok := s.commands[strings.Trim(cmd[0], "/")]
	if ok {
		handler(msg, &s)
		return
	}
	_, err := fmt.Fprintf(msg.src, "Huh, what?\n")
	if err != nil {
		log.Fatal(err)
	}
}

func newSimpleServer(c config) *simpleServer {
	addr := &net.TCPAddr{
		IP:   net.ParseIP(c.ip),
		Port: c.port,
	}

	commands := make(map[string]commandHandler)
	commands["who"] = whoCmdHandler
	commands["uptime"] = uptimeCmdHandler
	commands["rooms"] = roomsCmdHandler
	commands["here"] = hereCmdHandler
	commands["say"] = sayCmdHandler
	commands["name"] = nameCmdHandler

	return &simpleServer{
		userlist: &userlist{},
		commands: commands,
		msgsChan: make(chan message),
		listener: listener{
			addr:     addr,
			newConns: make(chan *net.Conn),
		},
		roomlist: []int{0},
	}
}

func (s *simpleServer) run() error {
	s.startTime = time.Now()
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
		if each.id != u.id {
			results = append(results, each)
		}
	}
	s.userlist.lock.Lock()
	s.userlist.users = results
	s.userlist.lock.Unlock()
}
