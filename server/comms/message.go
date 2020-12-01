package comms

import "github.com/crucialcarl/simpleclientserver/server/user"

type message struct {
	src user.User
	txt string
}
