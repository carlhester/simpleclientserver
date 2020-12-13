package main

import (
	"fmt"
	"log"
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
