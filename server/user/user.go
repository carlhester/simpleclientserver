package user

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// User ...
type User struct {
	Id         int
	Conn       clientFrontEnd
	Name       string
	Inbox  chan string // Msgs destined for this User
	Outbox chan string // Msgs from this user
	PList      UserList
	comm       communicator
	location   int
}

func (p *User) Close(msg string) {
	log.Printf("%s\n", msg)
	delete(p.PList, p.Id)
	p.Conn.Close()
}

func SetupNewUser(conn net.Conn, id int, UserList UserList, comm communicator) {
	var p *User
	var msgs = make(chan string)
	p = &User{
		Id:        id,
		Conn:      conn,
		Inbox: msgs,
		PList:     UserList,
		comm:      comm,
	}
	UserList[p.Id] = p
	go p.comm.ReceiveIncomingMessages(*p)
	go p.receiveMsgs()
	p.comm.SendMsgTo(fmt.Sprintf("You are connection ID: %d", id), p)
}

func (p *User) receiveMsgs() {
	for {
		txt, ok := <-p.Inbox
		if ok {
			writer := bufio.NewWriter(p.Conn)
			_, err := writer.WriteString(txt)
			if err != nil {
				p.Close(err.Error())
			}
			err = writer.Flush()
			if err != nil {
				p.Close(err.Error())
			}

		}
	}
}
