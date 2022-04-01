package main

import (
	"net/http"

	"github.com/takkiiiiiiiii/rest-api/controller"
	"github.com/takkiiiiiiiii/rest-api/model/repository"
)

//DI
var tr = repository.NewApiRepository()
var tc = controller.NewApiController(tr)
var ro = controller.NewRouter(tc)

func main() {
	http.HandleFunc("/api/users/", ro.HandleApiRequest)
	http.ListenAndServe(":7777", nil)
}
