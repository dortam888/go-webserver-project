package main

import (
	"net/http"

	"github.com/go_cli/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
