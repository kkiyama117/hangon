package usecases

import (
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
