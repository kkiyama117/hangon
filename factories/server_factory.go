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
)

type server struct {
	router chi.Router
}

type Server interface {
	Run(string) error
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

func InjectURIWithRoot(server *server) {
	// root
	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("root."))
	})

	// test panic
	server.router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})
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
	InjectURIWithRoot(server)
	// users
	server.router.Mount("/users", UserRouter())
	return server
}

func (server *server) Run(addr string) error {
	// return error
	return http.ListenAndServe(addr, server.router)
}
