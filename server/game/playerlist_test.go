package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerList_Add_First(t *testing.T) {
	// arrange
	p := playerList{}
	newPlayer := player{
		id: 123456,
	}

	// act
	p.Add(newPlayer)
	actual := p.players

	// assert
	expected := []player{player{id: 123456}}
	assert.Equal(t, expected, actual)
}

func TestPlayerList_Add_additional(t *testing.T) {
	// arrange
	p := playerList{
		players: []player{
			player{
				id: 123456,
			},
		},
	}
	newPlayer := player{
		id: 98765,
	}

	// act
	p.Add(newPlayer)
	actual := p.players

	// assert
	expected := []player{player{id: 123456}, player{id: 98765}}
	assert.Equal(t, expected, actual)
}

func TestPlayerList_Remove_Last(t *testing.T) {
	// arrange
	p := playerList{
		players: []player{
			player{
				id: 123456,
			},
		},
	}

	// act
	p.Remove(player{id: 123456})

	// assert
	assert.Equal(t, 0, len(p.players))
}

func TestPlayerList_Remove_Others(t *testing.T) {
	// arrange
	p := playerList{}
	p.Add(player{id: 55555})
	p.Add(player{id: 12345})
	p.Add(player{id: 11111})

	// act
	p.Remove(player{id: 12345})

	// assert
	expected := playerList{
		players: []player{
			player{id: 55555},
			player{id: 11111},
		},
	}
	assert.Equal(t, expected, p)
	assert.Equal(t, 2, len(p.players))
}

func TestPlayerList_Remove_NotFound(t *testing.T) {
	// arrange
	p := playerList{}
	p.Add(player{id: 55555})
	p.Add(player{id: 11111})

	// act
	p.Remove(player{id: 12345})

	// assert
	expected := playerList{
		players: []player{
			player{id: 55555},
			player{id: 11111},
		},
	}
	assert.Equal(t, expected, p)
	assert.Equal(t, 2, len(p.players))
}
