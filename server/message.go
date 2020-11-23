package main

import (
	"bufio"
	"fmt"
	"log"
)

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
			//log.Println(prefix + scanner.Text() + "\n")
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
