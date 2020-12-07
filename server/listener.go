package main

import "net"

type listener struct {
	newConns chan *net.Conn
	addr     *net.TCPAddr
}
