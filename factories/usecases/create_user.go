package usecases

import (
	"pumpkin/interface_adapters/create_user/controllers"
	"pumpkin/interface_adapters/create_user/driver_ports"
	"pumpkin/interface_adapters/create_user/presenters"
	"pumpkin/usecases/create_user/interactor"
)


func InjectedCreateUser(output driver_ports.APIOutput) controllers.CreateUserController{
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
