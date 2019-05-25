package factories

import (
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	. "pumpkin/codes/factories/handlers"
)

func UserRouter(dbFunc func()*gorm.DB) chi.Router {
	// users
	r := chi.NewRouter()
	r.Post("/", InjectCreateUser(dbFunc))
	r.Route("/{userID}", func(r chi.Router) {
		r.Get("/", nil)
	})
	return r
}
