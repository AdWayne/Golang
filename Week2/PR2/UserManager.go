package main

type UserManager struct {
	users []User
}

func (um *UserManager) AddUser(user User) {
	um.users = append(um.users, user)
}

func (um *UserManager) RemoveUser(email string) {
	index := um.findUserIndex(email)
	if index != -1 {
		um.users = append(um.users[:index], um.users[index+1:]...)
	}
}

func (um *UserManager) UpdateUser(email string, newUser User) {
	index := um.findUserIndex(email)
	if index != -1 {
		um.users[index] = newUser
	}
}

func (um *UserManager) findUserIndex(email string) int {
	for i, user := range um.users {
		if user.Email == email {
			return i
		}
	}
	return -1
}
