/*
Handler Factory

Handling usecase to http.handler
*/
package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"pumpkin/domain/model"
	"pumpkin/external_interfaces/server/view"
	"pumpkin/interface_adapters/create_user/controllers"
	"pumpkin/interface_adapters/create_user/presenters"
	"pumpkin/usecases/create_user/interactor"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Request
	// get data from request
	body, err := ioutil.ReadAll(r.Body)
	// get user data
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user.UserName ==""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Response
	w.Header().Set("Content-Type", "application/json")
	// DIPを鑑みて, Callback にする
	output := view.NewViewOutput(w)
	// push
	c := controllers.NewController(interactor.NewInteractor(presenters.NewPresenter(output)))
	// この内部でcallback が呼ばれて応答をする.
	err = c.CreateUser(&user)
}
