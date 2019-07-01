package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"pumpkin/codes/domain/model"
	"pumpkin/codes/factories/usecases/db"
	"pumpkin/codes/factories/usecases/html"
	"pumpkin/codes/framework_drivers"
)

func InjectGetUsers(dbFunc func() *gorm.DB) http.HandlerFunc {
	// usecase を 構築してから外部のinterfaces に渡す事で依存性の原則を満たす
	// handlerとしてサーバー(router)にAttachする起点の関数
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		// store user usecase
		var users model.Users
		// usecase を構築する.
		dbOutput := framework_drivers.NewDBOutput(dbFunc)
		getUsers := db.InjectedGetUsers(dbOutput)
		// Usecase処理実行
		err := getUsers.GetUsers(users)

		// Response
		w.Header().Set("Content-Type", "application/json")
		// DIP を意識して, callback の形でレスポンスの処理を埋め込ませる.
		output := framework_drivers.NewAPIOutput(w)
		// usecase を構築する.
		c := html.InjectedShowUsers(output)
		// 下の関数の内部でUsecaseの処理と injectorWithOutput が呼ばれて応答をする.
		err = c.ShowUsers(users)
		if err != nil {
			print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
