package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

type user struct {
	net.Conn
	id        int
	name      string
	msgsChan  chan message
	loginTime time.Time
	room      int
}

type userlist struct {
	lock  sync.Mutex
	users []*user
}

func newUser(id int, conn *net.Conn, msgsChan chan message) *user {
	return &user{
		Conn:      *conn,
		id:        id,
		name:      fmt.Sprintf("User-%d", id),
		msgsChan:  msgsChan,
		loginTime: time.Now(),
		room:      0,
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
