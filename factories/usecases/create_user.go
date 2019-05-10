package usecases

import (
	"pumpkin/external_interfaces"
	"pumpkin/interface_adapters/create_user/controllers"
	"pumpkin/interface_adapters/create_user/presenters"
	"pumpkin/usecases/create_user/interactor"
)


func InjectedCreateUser(output external_interfaces.Output) controllers.CreateUserController{
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
