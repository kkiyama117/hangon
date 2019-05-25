/*
Server Factory

Inject some parts and create web instance
*/
package factories

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"pumpkin/codes"
)

type server struct {
	router   chi.Router
	database gorm.DB
	config   codes.Configs
}

type Server interface {
	// call http.ListenAndServe
	Run(string) error
}

func NewServer() Server {
	// create instance
	sv := &server{}
	// initialize
	sv = Inject(sv)
	// return injected web server
	return sv
}

// inject router and handler and usecase
func Inject(server *server) *server {
	config := codes.GetConfigs()
	server.router = chi.NewRouter()
	db:= InjectDBFunc(WithDBOption(config))
	// Initialize web
	err := InjectMiddleware(server)
	if err != nil {
		panic("error with inject middleware")
	}
	// root router
	InjectURIWithRoot(server)
	// users
	server.router.Mount("/users", UserRouter(db))
	return server
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

	// get config environment
	server.router.Get("/config", func(w http.ResponseWriter, r *http.Request) {
		confs := codes.GetConfigs()
		for _, conf := range *confs {
			data, _ := json.Marshal(conf)
			_, _ = w.Write([]byte(data))
		}
	})

}

func (server *server) Run(addr string) error {
	// return error
	return http.ListenAndServe(addr, server.router)
}
