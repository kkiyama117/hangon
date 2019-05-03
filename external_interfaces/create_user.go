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
	"pumpkin/external_interfaces/common"
	"pumpkin/external_interfaces/web"
	"pumpkin/usecases/create_user/inputport"
)

func CreateUser(callback func(common.Output) inputport.CreateUserInputPort) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		if user.UserName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Response
		w.Header().Set("Content-Type", "application/json")
		// DIP を意識してcallback の形でレスポンスの処理を埋め込ませる.
		output := web.NewOutput(w)
		// usecase を構築する.
		c := callback(output)
		//下の関数の内部でUsecaseの処理と callback が呼ばれて応答をする.
		err = c.DoMethod(&user)
	}
}

