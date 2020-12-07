package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type user struct {
	net.Conn
	id      int
	name    string
	msgsChan chan message
}

type userlist struct {
	lock  sync.Mutex
	users []*user
}

func newUser(id int, conn *net.Conn, msgsChan chan message) *user {
	return &user{
		Conn:    *conn,
		id:      id,
		name:    fmt.Sprintf("User-%d", id),
		msgsChan: msgsChan,
	}
}

func (u user) process() {
	defer u.Close()
	for {
		scanner := bufio.NewScanner(u.Conn)
		for scanner.Scan() {
			txt := scanner.Text()
			msg := message{
				src: u,
				txt: txt,
			}
			u.msgsChan <- msg
		}
	}
}
