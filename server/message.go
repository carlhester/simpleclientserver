package main

import (
	"bufio"
	"fmt"
)

func sendMsgTo(errChan chan<- clientErr, msg string, players ...player) {
	if msg[len(msg)-1] != '\n' {
		msg = msg + string('\n')
	}
	for _, p := range players {
		writer := bufio.NewWriter(p.conn)
		_, err := writer.WriteString(msg)
		err = writer.Flush()
		if err != nil {
			newErr := clientErr{
				p:   &p,
				err: err,
			}
			errChan <- newErr
		}
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

func echoMessages(errChan chan<- clientErr, player player, players *playerList) {
	for {
		txt, ok := <-player.msgs
		if ok {
			for _, p := range players.players {
				if p != player {
					sendMsgTo(errChan, txt, p)
				}
			}
		}
	}
}
