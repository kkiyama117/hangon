package handlers

import (
	"net/http"

	"pumpkin/external_interfaces"
	"pumpkin/external_interfaces/common"
	"pumpkin/interface_adapters/create_user/controllers"
	"pumpkin/interface_adapters/create_user/presenters"
	"pumpkin/usecases/create_user/inputport"
	"pumpkin/usecases/create_user/interactor"
)

func InjectedUsecase(output common.Output) inputport.CreateUserInputPort {
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	external_interfaces.CreateUser(InjectedUsecase)(w, r)
}
