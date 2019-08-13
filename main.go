package main

import (
	"fmt"

	"github.com/jimxshaw/gowebservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Thomas",
		LastName:  "Jefferson",
	}

	fmt.Println(u)
}
