package main

import (
	"log"
	"net"
)

type listener struct {
	addr *net.TCPAddr
}

func (l listener) Listen(conns chan<- *net.Conn) {
	log.Printf("listening on %s", l.addr)
	listener, err := net.ListenTCP("tcp", l.addr)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	// accept network connections and assign players
	for {
		conn, err := listener.Accept()
		log.Println("Accepted")
		if err != nil {
			log.Panic(err)
		}
		conns <- &conn
	}

}
