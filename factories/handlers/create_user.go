package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"pumpkin/domain/model"
	"pumpkin/external_interfaces"
	us "pumpkin/factories/usecases"
)

// usecase を 構築してから外部のinterfaces に渡す事で依存性の原則を満たす
// handlerとしてサーバー(router)にAttachしないといけないのでその起点の関数
func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	output := external_interfaces.NewHttpOutput(w)
	// usecase を構築する.
	c := us.InjectedCreateUser(output)
	//下の関数の内部でUsecaseの処理と injectorWithOutput が呼ばれて応答をする.
	err = c.CreateUser(&user)
}
