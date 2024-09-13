package api

import (
	"be-skillacademy-final/service"

	"fmt"
	"net/http"
)

type API struct {
	userService    service.UserService
	sessionService service.SessionService
	mux            *http.ServeMux
}

func NewAPI(userService service.UserService, sessionService service.SessionService) API {
	mux := http.NewServeMux()
	api := API{
		userService,
		sessionService,
		mux,
	}

	mux.Handle("/user/register", http.HandlerFunc(api.Register))
	mux.Handle("/user/login", http.HandlerFunc(api.Login))
	mux.Handle("/user/logout", http.HandlerFunc(api.Logout))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}