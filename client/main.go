package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connect("0.0.0.0", 8123)

}

func connect(host string, port int) {
	ip := &net.TCPAddr{IP: net.ParseIP(host)}
	portNum := &net.TCPAddr{Port: port}
	conn, err := net.DialTCP("tcp", ip, portNum)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go receiveFromServer(conn)
	go receiveInput(conn)
	for {

	}
}

func receiveInput(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(conn)
	for scanner.Scan() {
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
	}
}

func receiveFromServer(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Print(scanner.Text() + "\n")
		}
	}
}
