/*
Handler Factory

Handling usecase to http.handler
*/
package external_interfaces

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"pumpkin/domain/model"
	"pumpkin/external_interfaces/web"
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
	// DIP を意識してcallback の形でレスポンスの処理を埋め込ませる.
	output := web.NewOutput(w)
	// usecase を構築する.
	c := controllers.NewController(interactor.NewInteractor(presenters.NewPresenter(output)))
	//下の関数の内部でUsecaseの処理と callback が呼ばれて応答をする.
	err = c.CreateUser(&user)
}
