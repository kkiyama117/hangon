package handlers

import (
	"net/http"

	"pumpkin/external_interfaces"
	us "pumpkin/factories/usecases"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// usecase を 構築してから外部のinterfaces に渡す事で依存性の原則を満たす
	// handlerとしてサーバー(router)にAttachしないといけないのでその起点の関数
	// external_interfaces内に直接handleFuncの形式(この関数の形式)で書いたほうが見易いが
	// factory(injector)と役割を明確に分離するためにファイルを分けた.
	// TODO: もっと上手い実装があるはず.
	external_interfaces.CreateUser(us.InjectedCreateUser)(w, r)
}
