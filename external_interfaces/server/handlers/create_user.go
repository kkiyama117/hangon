/*
Handler Factory

Handling usecase to http.handler
*/
package handlers

import (
	"encoding/json"
	"net/http"
	"pumpkin/domain/model"
	"pumpkin/external_interfaces/server/view"
	"pumpkin/interface_adapters/create_user/controllers"
	"pumpkin/interface_adapters/create_user/presenters"
	"pumpkin/usecases/create_user/interactor"
)
func GetHandlers(){}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Request
	// get data from request
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&model.User{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Response
	// DIPを鑑みて, Callback にする
	output := view.NewViewOutput(w)
	// push
	c := controllers.NewController(interactor.NewInteractor(presenters.NewPresenter(output)))
	err = c.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
