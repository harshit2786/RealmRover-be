package router

import (
	"realmrovers/handler"

	"github.com/gorilla/mux"
)

func NewRouter(u *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users",  u.GetUserById ).Methods("GET")
	return r 
}