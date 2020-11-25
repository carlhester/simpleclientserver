package server

import "github.com/crucialcarl/simpleclientserver/server/game"

type Server struct {
}

func (s Server) Run() {
	g := &game.Game{}
	g.Run()
}
