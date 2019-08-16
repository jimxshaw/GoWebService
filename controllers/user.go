package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// newUserController is a constructor function that
// returns a newly configured user controller.
func newUserController() *userController {
	// It is permissible to immediately use an address-of
	// operator on newly created structs.
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

// ServeHTTP takes in a request and returns a response.
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}
