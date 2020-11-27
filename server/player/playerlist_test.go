package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerList_Add_First(t *testing.T) {
	// arrange
	p := PlayerList{}
	newPlayer := Player{
		Id: 123456,
	}

	// act
	p.Add(newPlayer)
	actual := p.Players

	// assert
	expected := []Player{Player{Id: 123456}}
	assert.Equal(t, expected, actual)
}

func TestPlayerList_Add_additional(t *testing.T) {
	// arrange
	p := PlayerList{
		Players: []Player{
			Player{
				Id: 123456,
			},
		},
	}
	newPlayer := Player{
		Id: 98765,
	}

	// act
	p.Add(newPlayer)
	actual := p.Players

	// assert
	expected := []Player{Player{Id: 123456}, Player{Id: 98765}}
	assert.Equal(t, expected, actual)
}

func TestPlayerList_Remove_Last(t *testing.T) {
	// arrange
	p := PlayerList{
		Players: []Player{
			Player{
				Id: 123456,
			},
		},
	}

	// act
	p.Remove(Player{Id: 123456})

	// assert
	assert.Equal(t, 0, len(p.Players))
}

func TestPlayerList_Remove_Others(t *testing.T) {
	// arrange
	p := PlayerList{}
	p.Add(Player{Id: 55555})
	p.Add(Player{Id: 12345})
	p.Add(Player{Id: 11111})

	// act
	p.Remove(Player{Id: 12345})

	// assert
	expected := PlayerList{
		Players: []Player{
			Player{Id: 55555},
			Player{Id: 11111},
		},
	}
	assert.Equal(t, expected, p)
	assert.Equal(t, 2, len(p.Players))
}

func TestPlayerList_Remove_NotFound(t *testing.T) {
	// arrange
	p := PlayerList{}
	p.Add(Player{Id: 55555})
	p.Add(Player{Id: 11111})

	// act
	p.Remove(Player{Id: 12345})

	// assert
	expected := PlayerList{
		Players: []Player{
			Player{Id: 55555},
			Player{Id: 11111},
		},
	}
	assert.Equal(t, expected, p)
	assert.Equal(t, 2, len(p.Players))
}
