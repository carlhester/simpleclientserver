package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_roomExists(t *testing.T) {
	var tests = []struct {
		name     string
		rooms    []int
		arg      int
		expected bool
	}{
		{"works with one", []int{1}, 1, true},
		{"works with many 1", []int{1, 2, 3}, 1, true},
		{"works with many 2", []int{1, 2, 3}, 2, true},
		{"works with no matches", []int{1, 2, 3}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := simpleServer{
				roomlist: tt.rooms,
			}
			actual := s.roomExists(tt.arg)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestServer_usersInRoom(t *testing.T) {
	testuser1 := &user{id: 1, room: 0}
	testuser2 := &user{id: 2, room: 0}
	testuser3 := &user{id: 3, room: 1}
	testuser4 := &user{id: 4, room: 2}

	userlist := &userlist{
		users: []*user{testuser1, testuser2, testuser3, testuser4},
	}
	s := simpleServer{
		userlist: userlist,
	}

	actual := s.usersInRoom(0)
	assert.Equal(t, []*user{testuser1, testuser2}, actual)
	actual = s.usersInRoom(1)
	assert.Equal(t, []*user{testuser3}, actual)
	actual = s.usersInRoom(2)
	assert.Equal(t, []*user{testuser4}, actual)
	actual = s.usersInRoom(99)
	assert.Equal(t, []*user{}, actual)
}
