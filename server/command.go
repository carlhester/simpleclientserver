package main

import (
	"fmt"
	"log"
)

type command interface {
	execute()
}

type whoCommand struct {
	msg      message
	userlist userlist
}

func (cmd whoCommand) execute() {
	result := fmt.Sprintf("NAME\tID\tTIME\n")
	for _, u := range cmd.userlist.users {
		result = result + fmt.Sprintf("%s\t%d\t%s\n", u.name, u.id, u.loginTime.Format("Mon Jan 2 15:04:05 MST 2006"))
		log.Print(result)
		_, err := fmt.Fprintf(cmd.msg.src, result)
		if err != nil {
			log.Fatal(err)
		}
	}

}
