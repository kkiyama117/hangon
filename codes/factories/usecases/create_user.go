package usecases

import (
	"pumpkin/codes/interface_adapters/create_user/controllers"
	"pumpkin/codes/interface_adapters/create_user/driver_ports"
	"pumpkin/codes/interface_adapters/create_user/presenters"
	"pumpkin/codes/usecases/create_user/interactor"
)

func InjectedCreateUser(output driver_ports.APIOutput) controllers.CreateUserController {
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
