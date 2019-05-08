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
	"pumpkin/interface_adapters/create_user/controllers"
)

func CreateUser(injectorWithOutput func(common.Output) controllers.CreateUserController) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		// get data from request
		body, err := ioutil.ReadAll(r.Body)
		// get user data
		user := model.User{}
		if err != nil {
			print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(body, &user)
		if user.UserName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// User model をゲットした

		// Response
		w.Header().Set("Content-Type", "application/json")
		// DIP を意識して, callback の形でレスポンスの処理を埋め込ませる.
		output := common.NewOutput(w)
		// usecase を構築する.
		c := injectorWithOutput(output)
		//下の関数の内部でUsecaseの処理と injectorWithOutput が呼ばれて応答をする.
		err = c.CreateUser(&user)
	}
}
