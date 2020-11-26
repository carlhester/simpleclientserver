package comms

import (
	"log"
	"net"
)

func Listen(addr *net.TCPAddr, conns chan<- *net.Conn) {
	log.Printf("listening on %s", addr)
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	// accept network connections and assign players
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		log.Printf("Client connected: %s...\n", conn.RemoteAddr())
		conns <- &conn
	}

}
