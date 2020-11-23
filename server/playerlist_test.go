package main

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
	p.add(newPlayer)
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
	p.add(newPlayer)
	actual := p.players

	// assert
	expected := []player{player{id: 123456}, player{id: 98765}}
	assert.Equal(t, expected, actual)
}
