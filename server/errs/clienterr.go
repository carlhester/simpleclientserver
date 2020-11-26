package errs

import "github.com/crucialcarl/simpleclientserver/server/player"

type ClientErr struct {
	P   *player.Player
	err error
}
