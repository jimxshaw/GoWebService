package main

import (
	"net/http"

	"github.com/jimxshaw/gowebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
