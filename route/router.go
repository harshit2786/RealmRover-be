package router

import (
	// "net/http"
	"net/http"
	"realmrovers/handler"

	// "realmrovers/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter(u *handler.UserHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/users", u.GetUserById).Methods("GET")
	r.HandleFunc("/github/token", u.SignUser).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:5173"},
		AllowCredentials: true,
	})
	// Create Map
	// Get Map
	handler := c.Handler(r)
	return handler
}
