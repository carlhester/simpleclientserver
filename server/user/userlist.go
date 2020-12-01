package user

type UserList map[int]*User

func NewUserList() map[int]*User {
	return make(map[int]*User)
}
