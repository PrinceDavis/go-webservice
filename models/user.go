package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func Adduser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("new user must not include an ID")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf(`User with ID '%v'  not found`, id)
}

func UpdateUser(user User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == user.ID {
			users[i] = &user
			return u, nil
		}
	}
	return User{}, fmt.Errorf(`User with ID '%v'  not found`, id)
}

func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1]...)
			return nil
		}
	}
	return User{}, fmt.Errorf(`User with ID '%v'  not found`, id)
}
