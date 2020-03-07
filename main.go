package main

import (
	"net/http"

	controller "github.com/princeDavis/webservice/controllers"
)

func main() {
	controller.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
