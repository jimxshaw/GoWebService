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

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}

	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteHeader([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
