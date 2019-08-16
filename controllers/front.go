package controllers

import "net/http"

// RegisterControllers will set up the routing
// in the entire application. When a network
// request is received, it will go to the correct
// controller to be processed.
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
