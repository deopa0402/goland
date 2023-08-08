package repository

import "main/domain"

var UserDB = map[int]domain.User{}

func SaveUser(User domain.User) domain.User {
	UserDB[User.ID] = User
	return User
}

func FindUserByID(id int) domain.User {
	return UserDB[id]
}
