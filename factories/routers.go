package factories

import (
	"github.com/go-chi/chi"
	"pumpkin/factories/handlers"
)

func UserRouter()chi.Router{
	// users
	r:=chi.NewRouter()
	r.Post("/", handlers.CreateUser)
	r.Route("/{userID}",func(r chi.Router){
		r.Get("/",nil)
	})
	return r
}
