package factories

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"pumpkin/codes"
	. "pumpkin/codes/factories/handlers"
)

func MainRouter() chi.Router {
	router := chi.NewRouter()
	// root
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("root."))
	})

	// test panic
	router.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// get config environment
	router.Get("/config", func(w http.ResponseWriter, r *http.Request) {
		pConf := codes.GetConfigs()
		cr := *pConf
		jsonData, _ := json.Marshal(cr)
		_, _ = w.Write([]byte(jsonData))
	})
	return router
}

func UserRouter(dbFunc func() *gorm.DB) chi.Router {
	// users
	r := chi.NewRouter()
	r.Get("/", InjectGetUsers(dbFunc))
	r.Post("/", InjectCreateUser(dbFunc))
	r.Route("/{userID}", func(r chi.Router) {
		r.Get("/", InjectGetUser(dbFunc))
	})
	return r
}
