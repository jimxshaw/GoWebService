package models

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
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
