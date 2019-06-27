package usecases

import (
	"pumpkin/codes/interface_adapters/show_user/controllers"
	"pumpkin/codes/interface_adapters/show_user/driver_ports"
	"pumpkin/codes/interface_adapters/show_user/presenters"
	"pumpkin/codes/usecases/_show_user/interactor"
)

func InjectedShowUser(output driver_ports.APIOutput) controllers.CreateUserController {
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
