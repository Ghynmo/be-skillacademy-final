package api

import (
	// "be-skillacademy-final/service"

	"fmt"
	"net/http"
)

type API struct {
	// userService    service.UserService
	mux            *http.ServeMux
}

func NewAPI() API { //userService service.UserService as parameter
	mux := http.NewServeMux()
	api := API{
		// userService,
		mux,
	}

	// mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}