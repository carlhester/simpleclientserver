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

func (p *playerList) remove(toRemove player) {
	var newList []player
	for _, p := range p.players {
		if p.id == toRemove.id {
			continue
			newList = append(newList, p)
		}
	}
	p.players = newList
}

// game handles the high level "global" state of the game
type game struct {
	playerList
}

// a player is a client with some labels
type player struct {
	id    int
	conn  clientFrontEnd
	name  string
	msgs  chan string
	pList *playerList
}

func main() {
	// initialize id incrementer
	id := make(chan int)
	go incrementer(id)

	// init empty playerList
	playerList := &playerList{}

	// the serverConsole uses standard in/out
	serverConsole := serverConsole{
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	}

	// players keeps a list of active players
	// the initial player is the serverConsole
	consoleId := <-id
	console := player{
		id:    consoleId,
		conn:  serverConsole,
		name:  "server",
		pList: playerList,
	}
	playerList.add(console)

	// new game
	g := &game{
		*playerList,
	}

	// create players by listening for network connects
	addr := &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: 8123}
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
		id := <-id
		go setupNewPlayer(conn, g, id, playerList)
	}
}

func setupNewPlayer(conn net.Conn, game *game, id int, playerList *playerList) {
	var newPlayer *player
	var msgs = make(chan string)
	log.Printf("Client connected: %s...\n", conn.RemoteAddr())
	newPlayer = &player{
		id:    id,
		conn:  conn,
		msgs:  msgs,
		pList: playerList,
	}
	getPlayerName(newPlayer)
	game.playerList.add(*newPlayer)
	sendMsgTo(fmt.Sprintf("You are player %d", id), *newPlayer)
	go listenForMessages(*newPlayer)
	go echoMessages(*newPlayer, &game.playerList)
}

func getPlayerName(p *player) {
	sendMsgTo("Hello! What is your name? ", *p)
	reader := bufio.NewReader(p.conn)
	name, _ := reader.ReadString('\n')
	p.name = name[:len(name)-1]
	log.Println(name)
}

func sendMsgTo(msg string, players ...player) {
	if msg[len(msg)-1] != '\n' {
		msg = msg + string('\n')
	}
	for _, v := range players {
		writer := bufio.NewWriter(v.conn)
		_, err := writer.WriteString(msg)
		if err != nil {
			log.Panic(err)
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
			for _, p := range players.players {
				if p != player {
					sendMsgTo(txt, p)
				}
			}
		}
	}
}

/*func ListenForConnection(addr *net.TCPAddr) *net.Conn {
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Client connected: %s...\n", conn.RemoteAddr())
	return &conn
}
*/
