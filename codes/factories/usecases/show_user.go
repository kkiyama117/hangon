package usecases

import (
	"pumpkin/codes/interface_adapters/_show_user/controllers"
	"pumpkin/codes/interface_adapters/_show_user/driver_ports"
	"pumpkin/codes/interface_adapters/_show_user/presenters"
	"pumpkin/codes/usecases/_show_user/interactor"
)

func InjectedShowUser(output driver_ports.APIOutput) controllers.ShowUserController {
	pres := presenters.NewPresenter(output)
	inter := interactor.NewInteractor(pres)
	c := controllers.NewController(inter)
	return c
}
