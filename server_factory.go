/*
Server Factory

Inject some parts and create server instance
*/
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

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
	server.router.Use(middleware.RequestID)
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Use(middleware.URLFormat)
	server.router.Use(render.SetContentType(render.ContentTypeJSON))

	// root
	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("root."))
	})

	// ping
	server.router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	// test panic
	server.router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})
	// users
	server.router.Post("/users", handlers.CreateUser)
	return server
}

func (server *server) Run() error {
	// return error
	return http.ListenAndServe(":3000", server.router)
}
