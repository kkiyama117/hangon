package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"pumpkin/codes/domain/model"
	"pumpkin/codes/factories/usecases/db"
	"pumpkin/codes/framework_drivers"
)

func InjectGetUser(dbFunc func() *gorm.DB) http.HandlerFunc {
	// usecase を 構築してから外部のinterfaces に渡す事で依存性の原則を満たす
	// handlerとしてサーバー(router)にAttachする起点の関数
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		// get data from request
		body, err := ioutil.ReadAll(r.Body)
		// get user data
		user := model.User{
			ID: 9,
			UserName: "test2",
			Mail: "test@test.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		}
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

		// store user usecase
		// usecase を構築する.
		// dbOutput := framework_drivers.NewDBOutput(dbFunc)
		// storeUser := usecases.InjectedStoreUser(dbOutput)
		// Usecase処理実行
		// err = storeUser.StoreUser(&user)
		if user.ID < 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Response
		w.Header().Set("Content-Type", "application/json")
		// DIP を意識して, callback の形でレスポンスの処理を埋め込ませる.
		output := framework_drivers.NewAPIOutput(w)
		// usecase を構築する.
		c := db.InjectedShowUser(output)
		// 下の関数の内部でUsecaseの処理と injectorWithOutput が呼ばれて応答をする.
		err = c.ShowUser(&user)
	}
}
