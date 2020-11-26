package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerList_Add_First(t *testing.T) {
	// arrange
	p := PlayerList{}
	newPlayer := Player{
		id: 123456,
	}

	// act
	p.Add(newPlayer)
	actual := p.players

	// assert
	expected := []Player{Player{id: 123456}}
	assert.Equal(t, expected, actual)
}

func TestPlayerList_Add_additional(t *testing.T) {
	// arrange
	p := PlayerList{
		players: []Player{
			Player{
				id: 123456,
			},
		},
	}
	newPlayer := Player{
		id: 98765,
	}

	// act
	p.Add(newPlayer)
	actual := p.players

	// assert
	expected := []Player{Player{id: 123456}, Player{id: 98765}}
	assert.Equal(t, expected, actual)
}

func TestPlayerList_Remove_Last(t *testing.T) {
	// arrange
	p := PlayerList{
		players: []Player{
			Player{
				id: 123456,
			},
		},
	}

	// act
	p.Remove(Player{id: 123456})

	// assert
	assert.Equal(t, 0, len(p.players))
}

func TestPlayerList_Remove_Others(t *testing.T) {
	// arrange
	p := PlayerList{}
	p.Add(Player{id: 55555})
	p.Add(Player{id: 12345})
	p.Add(Player{id: 11111})

	// act
	p.Remove(Player{id: 12345})

	// assert
	expected := PlayerList{
		players: []Player{
			Player{id: 55555},
			Player{id: 11111},
		},
	}
	assert.Equal(t, expected, p)
	assert.Equal(t, 2, len(p.players))
}

func TestPlayerList_Remove_NotFound(t *testing.T) {
	// arrange
	p := PlayerList{}
	p.Add(Player{id: 55555})
	p.Add(Player{id: 11111})

	// act
	p.Remove(Player{id: 12345})

	// assert
	expected := PlayerList{
		players: []Player{
			Player{id: 55555},
			Player{id: 11111},
		},
	}
	assert.Equal(t, expected, p)
	assert.Equal(t, 2, len(p.players))
}
