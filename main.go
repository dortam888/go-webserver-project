package main

import (
	"net/http"

	"github.com/go_cli/controllers"
)

func main() {

	//listen to web server in port 3000
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
