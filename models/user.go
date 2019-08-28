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

// GetUserByID returns the requested user.
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			// Get user by dereferencing that pointer because
			// a user value is expected.
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser sets the user id.
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			// Use the index to replace the entry in the users
			// slice with the user that was provided.
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID deletes specified user.
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			// Build a new slice with the portion before the found
			// user and the portion after the found user.
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
