package player

type PlayerList struct {
	players []Player
}

func (p *PlayerList) Add(newplayer Player) {
	p.players = append(p.players, newplayer)
}

func (p *PlayerList) Remove(toRemove Player) {
	var newList []Player
	for _, p := range p.players {
		if p.id == toRemove.id {
			continue
		}
		newList = append(newList, p)
	}
	p.players = newList
}

func (p PlayerList) Get() []Player {
	return p.players
}
