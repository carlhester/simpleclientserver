package player

type PlayerList struct {
	Players []Player
}

func (p *PlayerList) Add(newplayer Player) {
	p.Players = append(p.Players, newplayer)
}

func (p *PlayerList) Remove(toRemove Player) {
	var newList []Player
	for _, p := range p.Players {
		if p.Id == toRemove.Id {
			continue
		}
		newList = append(newList, p)
	}
	p.Players = newList
}

func (p PlayerList) Get() []Player {
	return p.Players
}
