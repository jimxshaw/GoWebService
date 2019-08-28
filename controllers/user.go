package controllers

import (
	"net/http"
	"regexp"

	"github.com/jimxshaw/gowebservice/models"
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

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeResponseAsJSON(u, w)
}
