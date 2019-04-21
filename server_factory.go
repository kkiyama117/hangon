/*
Server Factory

Inject some parts and create server instance
*/
package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"pumpkin/external_interfaces/server/handlers"
)

type server struct {
	router chi.Router
}

type Server interface {
	Run() error
}

func NewServer() Server {
	// create instance
	sv := &server{}
	// initialize
	sv = Inject(sv)
	// return injected server
	return sv
}

// inject router and handler and usecase
func Inject(server *server) *server {
	server.router = chi.NewRouter()
	// Initialize server
	server.router.Post("/users", handlers.CreateUser)
	return server
}

func (server *server) Run() error {
	// return error
	return http.ListenAndServe(":3000", server.router)
}
