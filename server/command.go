package main

import (
	"fmt"
	"log"
	"time"
)

type commandHandler func(message, *simpleServer)

func whoCmdHandler(msg message, s *simpleServer) {
	output := fmt.Sprintf("NAME\tID\tTIME\n")
	for _, u := range s.userlist.users {
		output = output + fmt.Sprintf("%s\t%d\t%s\n", u.name, u.id, u.loginTime.Format("Mon Jan 2 15:04:05 MST 2006"))
	}
	_, err := fmt.Fprintf(msg.src, output)
	if err != nil {
		log.Fatal(err)
	}
}

func roomsCmdHandler(msg message, s *simpleServer) {
	output := fmt.Sprintf("ID\n")
	for _, r := range s.roomlist {
		output = output + fmt.Sprintf("%d\n", r)
	}
	_, err := fmt.Fprintf(msg.src, output)
	if err != nil {
		log.Fatal(err)
	}
}

func hereCmdHandler(msg message, s *simpleServer) {
	output := "You look around and see "
	usersHere := s.usersInRoom(msg.src.room)

	for i, u := range usersHere {
		if len(usersHere) == i+1 {
			output = output + fmt.Sprintf("%s", u.name)
		} else {
			output = output + fmt.Sprintf("%s and ", u.name)
		}
	}
	output = output + fmt.Sprintf(" here.\n")
	_, err := fmt.Fprintf(msg.src, output)
	if err != nil {
		log.Fatal(err)
	}
}

func uptimeCmdHandler(msg message, s *simpleServer) {
	output := fmt.Sprintf("%s\n", time.Since(s.startTime).Truncate(time.Second).String())
	_, err := fmt.Fprintf(msg.src, output)
	if err != nil {
		log.Fatal(err)
	}
}
