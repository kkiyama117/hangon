package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"pumpkin/codes/domain/model"
	"pumpkin/codes/factories/usecases/db"
	"pumpkin/codes/factories/usecases/html"
	"pumpkin/codes/framework_drivers"
)

func InjectGetUser(dbFunc func() *gorm.DB) http.HandlerFunc {
	// usecase を 構築してから外部のinterfaces に渡す事で依存性の原則を満たす
	// handlerとしてサーバー(router)にAttachする起点の関数
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		// get data from request
		// see router.go in parent folder
		userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
		// get user data
		user := model.User{
			ID:        uint(userID),
			DeletedAt: nil,
		}
		// show user usecase
		// usecase を構築する.
		dbOutput := framework_drivers.NewDBOutput(dbFunc)
		getUser := db.InjectedGetUser(dbOutput)
		// Usecase処理実行
		err := getUser.GetUser(&user)
		if user.ID < 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Response
		w.Header().Set("Content-Type", "application/json")
		// DIP を意識して, callback の形でレスポンスの処理を埋め込ませる.
		output := framework_drivers.NewAPIOutput(w)
		// usecase を構築する.
		c := html.InjectedShowUser(output)
		// 下の関数の内部でUsecaseの処理と injectorWithOutput が呼ばれて応答をする.
		err = c.ShowUser(&user)

		// cache error
		if err != nil {
			print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
