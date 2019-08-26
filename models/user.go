package models

import (
	"errors"
	"fmt"
)

// User is the default user object.
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers will retrieve a slice of users.
func GetUsers() []*User {
	return users
}

// AddUser adds a user to the users slice.
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserById returns the requested user.
func GetUserById(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			// Get user by dereferencing that pointer because
			// a user value is expected.
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}
