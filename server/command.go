package main

import (
	"fmt"
	"log"
)

type command struct {
	directive string
	msg       message
	state     *simpleServer
}

func (cmd command) execute() {
	switch cmd.directive {
	case "who":
		result := fmt.Sprintf("NAME\tID\tTIME\n")
		for _, u := range cmd.state.userlist.users {
			result = result + fmt.Sprintf("%s\t%d\t%s\n", u.name, u.id, u.loginTime.Format("Mon Jan 2 15:04:05 MST 2006"))
			log.Print(result)
			_, err := fmt.Fprintf(cmd.msg.src, result)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
