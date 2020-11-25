package game

type PlayerList struct {
	players []player
}

func (p *PlayerList) Add(newplayer player) {
	p.players = append(p.players, newplayer)
}

func (p *PlayerList) Remove(toRemove player) {
	var newList []player
	for _, p := range p.players {
		if p.id == toRemove.id {
			continue
		}
		newList = append(newList, p)
	}
	p.players = newList
}

func (p PlayerList) Get() []player {
	return p.players
}
