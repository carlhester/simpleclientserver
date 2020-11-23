package main

type playerList struct {
	players []player
}

func (p *playerList) add(player player) {
	p.players = append(p.players, player)
}

func (p *playerList) remove(toRemove player) {
	var newList []player
	for _, p := range p.players {
		if p.id == toRemove.id {
			continue
		}
		newList = append(newList, p)
	}
	p.players = newList
}

func (p playerList) Get() []player {
	return p.players
}
