/*
Server Factory

Inject some parts and create web instance
*/
package factories

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"pumpkin/external_interfaces"
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
	// return injected web
	return sv
}

func InjectMiddleware(server *server) error {
	if server.router == nil {
		return errors.New("nil web, please set web instance")
	}
	server.router.Use(middleware.RedirectSlashes)
	server.router.Use(middleware.RequestID)
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	// for ping
	server.router.Use(middleware.Heartbeat("/ping"))
	server.router.Use(middleware.URLFormat)
	server.router.Use(render.SetContentType(render.ContentTypeJSON))
	return nil
}

func InjectURIWithRoot(server *server) error {
	// root
	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("root."))
	})

	// test panic
	server.router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})
	return nil
}

// inject router and handler and usecase
func Inject(server *server) *server {
	server.router = chi.NewRouter()
	// Initialize web
	err := InjectMiddleware(server)
	if err != nil {
		panic("error with inject middleware")
	}
	// root router
	err = InjectURIWithRoot(server)
	if err != nil {
		panic("error with inject root router")
	}
	// users
	server.router.Post("/users", external_interfaces.CreateUser)
	return server
}

func (server *server) Run() error {
	// return error
	return http.ListenAndServe(":3000", server.router)
}