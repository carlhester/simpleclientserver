package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// clientFrontEnd is how the server interacts with the external clients
type clientFrontEnd interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
}

// the serverConsole is a special kind of client
type serverConsole struct {
	*bufio.Reader
	*bufio.Writer
}

type playerList struct {
	players []player
}

func (p *playerList) add(player player) {
	p.players = append(p.players, player)
}

// game handles the high level "global" state of the game
type game struct {
	playerList
}

// a player is a client with some labels
type player struct {
	id   int
	conn clientFrontEnd
	name string
	msgs chan string
}

func main() {
	// the serverConsole uses standard in/out
	serverConsole := serverConsole{
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	}

	// players keeps a list of active players
	// the initial player is the serverConsole
	console := player{
		id:   0,
		conn: serverConsole,
		name: "server",
	}
	playerList := &playerList{}
	playerList.add(console)

	// new game
	g := game{
		*playerList,
	}

	// create players by listening for network connects
	addr := &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123}
	log.Printf("listening on %s", addr)
	// there are 3 players: server, p1 and p2
	for id := 1; id < 3; id++ {
		conn := ListenForConnection(addr)
		msgs := make(chan string)
		player := &player{
			id:   id,
			conn: *conn,
			msgs: msgs,
		}
		getPlayerName(player)
		g.playerList.add(*player)

		sendMsgTo(fmt.Sprintf("You are player %d", id), *player)
		go listenForMessages(*player)
		go echoMessages(*player, &g.playerList)
	}
	log.Println(g)

	// test writing to each
	sendMsgTo("Server: BROADCAST", g.playerList.players...)
	for _, v := range g.playerList.players {
		scanner := bufio.NewScanner(v.conn)
		if scanner.Scan() {
			log.Println(scanner.Text())
		}
	}
}

func getPlayerName(p *player) {
	sendMsgTo("Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.conn)
	name, _ := reader.ReadString('\n')
	p.name = name[:len(name)-1]
	log.Println(name)
}

func sendMsgTo(msg string, players ...player) {
	// Remove newline if exists and add our own
	if msg[len(msg)-1] != '\n' {
		msg = msg + string('\n')
	}
	for _, v := range players {
		writer := bufio.NewWriter(v.conn)
		_, err := writer.WriteString(msg)
		if err != nil {
			log.Println(err)
		}
		writer.Flush()
	}
}

func listenForMessages(p player) {
	for {
		scanner := bufio.NewScanner(p.conn)
		for scanner.Scan() {
			prefix := fmt.Sprintf("[%d] %s: ", p.id, p.name)
			log.Println(prefix + scanner.Text() + "\n")
			p.msgs <- prefix + scanner.Text() + "\n"
		}
	}
}

func echoMessages(player player, players *playerList) {
	for {
		txt, ok := <-player.msgs
		if ok {
			fmt.Println("ok!")
			sendMsgTo(txt, players.players...)
			log.Println(players)
		}
	}
}

func ListenForConnection(addr *net.TCPAddr) *net.Conn {
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Client connected: %s...\n", conn.RemoteAddr())
	return &conn
}
